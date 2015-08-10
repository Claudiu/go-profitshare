[![](https://camo.githubusercontent.com/bfdd3541106bf567a1c4339af2cbf33fc60257e2/68747470733a2f2f662e636c6f75642e6769746875622e636f6d2f6173736574732f343536362f313133353630352f61623439323939302d316331392d313165332d383633622d6466646337653635313766312e706e67)](https://godoc.org/github.com/Claudiu/go-profitshare)


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
    		"api user",
    		"api key",
    	)

    	advertisers, err := ps.GetAdvertisers1M()

      if err != nil {
        panic(err)
      }

    	for _, advertiser := range advertisers {
    		fmt.Println(advertiser.ID)
    	}
    }
