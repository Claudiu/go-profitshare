package profitshare

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Commissions instance
type Commissions struct {
	CommissionType string `json:"type"`
	Value          string `json:"value"`
}

// Advertiser instance
type Advertiser struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	Logo               string        `json:"logo"`
	Category           string        `json:"category"`
	URL                string        `json:"url"`
	LastUpdateProducts string        `json:"last_update_products"`
	Commissions        []Commissions `json:"commissions"`
}

// GetAdvertisers returns a list of advertise from a period of time to another
func (ps *ProfitShare) GetAdvertisers(from time.Time, to time.Time) []Advertiser {
	dateFormat := "2006-01-02"

	url, err := url.Parse("affiliate-advertisers")

	if err != nil {
		fmt.Println(err)
	}

	q := url.Query()
	q.Add("date_from", from.Format(dateFormat))
	q.Add("date_to", to.Format(dateFormat))
	url.RawQuery = q.Encode()

	body := ps.Get(url.String())

	var rez map[string]map[string]Advertiser
	_ = json.Unmarshal(body, &rez)

	var ret []Advertiser
	for _, item := range rez["result"] {
		ret = append(ret, item)
	}

	return ret
}

// GetAdvertisers1M returns a list of advertisers from a month ago to now
func (ps *ProfitShare) GetAdvertisers1M() []Advertiser {
	now := time.Now()
	return ps.GetAdvertisers(now.AddDate(0, -1, 0), now)
}

// GetAdvertisers1D returns a list of advertisers from a day ago to now
func (ps *ProfitShare) GetAdvertisers1D() []Advertiser {
	now := time.Now()
	return ps.GetAdvertisers(now.AddDate(0, 0, -1), now)
}
