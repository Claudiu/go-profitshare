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
func (ps *ProfitShare) GetCampaignPage(page int) ([]Campaign, Paginator) {
	url, err := url.Parse("affiliate-campaigns")

	if err != nil {
		fmt.Println(err)
	}

	q := url.Query()
	q.Add("page", fmt.Sprintf("%d", page))
	url.RawQuery = q.Encode()

	body := ps.Get(url.String())

	rez := map[string]CampaignsResult{}
	_ = json.Unmarshal(body, &rez)

	return rez["result"].Campaigns, rez["result"].Paginator
}

// GetCampaigns returns all the active campaigns
func (ps *ProfitShare) GetCampaigns() []Campaign {
	campaigns, pag := ps.GetCampaignPage(1)

	for index := 2; index <= pag.TotalPages; index++ {
		currentCampaign, _ := ps.GetCampaignPage(index)
		campaigns = append(campaigns, currentCampaign...)
		time.Sleep(ps.SleepTime)
	}

	return campaigns
}
