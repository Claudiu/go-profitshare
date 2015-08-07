![https://godoc.org/github.com/Claudiu/go-profitshare](https://camo.githubusercontent.com/bfdd3541106bf567a1c4339af2cbf33fc60257e2/68747470733a2f2f662e636c6f75642e6769746875622e636f6d2f6173736574732f343536362f313133353630352f61623439323939302d316331392d313165332d383633622d6466646337653635313766312e706e67)


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
