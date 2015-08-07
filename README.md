Currently Implemented:
- Grabbing the list of Advertisers with Commissions
- Grabbing current Campaigns with Banners
- Grabbing Products with Advertiser ID filtering


Grabbing the list of advertisers
--------------------------------
    package main
    
    import (
    	"fmt"
    	"github.com/claudiu/go-profitshare"
    )
    
    func main() {
    	ps := profitshare.NewProfitShare(
    		"claudiu",                          // User
    		"dqwdwqdd1r32r2332f23f3f23f32fsad", // API Key
    	)
    
    	advertisers := ps.GetAdvertisers1M()
    
    	for _, advertiser := range advertisers {
    		fmt.Println(advertiser.ID)
    	}
    }
