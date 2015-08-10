package profitshare

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Banner holds Campaign banner information
type Banner struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Src    string `json:"src"`
}

// Campaign holds Campaign information
type Campaign struct {
	ID             int               `json:"id"`
	AdvertiserID   int               `json:"advertiser_id"`
	Name           string            `json:"name"`
	CommissionType string            `json:"commissionType"`
	StartDate      string            `json:"startDate"`
	EndDate        string            `json:"endDate"`
	PageURL        string            `json:"url"`
	Banners        map[string]Banner `json:"banners"`
}

// CampaignsResult holds the result of the affiliate-campaigns API call
type CampaignsResult struct {
	Paginator Paginator  `json:"paginator"`
	Campaigns []Campaign `json:"campaigns"`
}

// GetCampaignPage returns all the campaigns from the supplied page
func (ps *ProfitShare) GetCampaignPage(page int) ([]Campaign, Paginator, error) {
	url, err := url.Parse("affiliate-campaigns")

	if err != nil {
		return []Campaign{}, Paginator{}, err
	}

	q := url.Query()
	q.Add("page", fmt.Sprintf("%d", page))
	url.RawQuery = q.Encode()

	body, err := ps.Get(url.String())

	if err != nil {
		return []Campaign{}, Paginator{}, err
	}

	rez := map[string]CampaignsResult{}
	err = json.Unmarshal(body, &rez)

	if err != nil {
		return []Campaign{}, Paginator{}, err
	}

	return rez["result"].Campaigns, rez["result"].Paginator, nil
}

// GetCampaigns returns all the active campaigns
func (ps *ProfitShare) GetCampaigns() ([]Campaign, error) {
	campaigns, pag, err := ps.GetCampaignPage(1)

	if err != nil {
		return []Campaign{}, err
	}

	for index := 2; index <= pag.TotalPages; index++ {
		currentCampaign, _, err := ps.GetCampaignPage(index)

		if err != nil {
			return []Campaign{}, err
		}

		campaigns = append(campaigns, currentCampaign...)
		time.Sleep(ps.SleepTime)
	}

	return campaigns, nil
}
