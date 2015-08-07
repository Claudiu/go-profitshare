package profitshare

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Product holds information about a certain product
type Product struct {
	AdvertiserID    int     `json:"advertiser_id"`
	AdvertiserName  string  `json:"advertiser_name"`
	BrandName       string  `json:"brand_name"`
	CategoryName    string  `json:"category_name"`
	Description     string  `json:"description"`
	FreeShipping    int     `json:"free_shipping"`
	Image           string  `json:"image"`
	ImageOriginal   string  `json:"image_original"`
	LastUpdate      string  `json:"last_update"`
	Link            string  `json:"link"`
	Name            string  `json:"name"`
	PartNumber      string  `json:"part_number"`
	Price           float64 `json:"price"`
	PriceDiscounted string  `json:"price_discounted"`
	PriceVat        float64 `json:"price_vat"`
}

// ProductsResult holds the result from the API
type ProductsResult struct {
	Result struct {
		CurrentPage    int       `json:"current_page"`
		RecordsPerPage int       `json:"records_per_page"`
		TotalPages     int       `json:"total_pages"`
		Products       []Product `json:"products"`
	} `json:"result"`
}

// GetProductPage returns a list of products and the paginator for a certain page
func (ps *ProfitShare) GetProductPage(advertiserID []int, page int) ([]Product, Paginator) {
	url, err := url.Parse("affiliate-products")

	if err != nil {
		fmt.Println(err)
	}

	str := make([]string, len(advertiserID))
	for index, item := range advertiserID {
		str[index] = fmt.Sprintf("%d", item)
	}

	q := url.Query()
	q.Add("page", fmt.Sprintf("%d", page))

	url.RawQuery = q.Encode()

	// Workaround: [ Encoding
	body := ps.Get(url.String() + "filters[advertisers]=" + strings.Join(str, ","))

	rez := ProductsResult{}
	_ = json.Unmarshal(body, &rez)

	return rez.Result.Products, Paginator{
		ItemsPerPage: rez.Result.RecordsPerPage,
		CurrentPage:  rez.Result.CurrentPage,
		TotalPages:   rez.Result.TotalPages,
	}
}

// GetProducts returns a list of all the products by advertiserIds
func (ps *ProfitShare) GetProducts(advertiserID []int) []Product {
	products, pag := ps.GetProductPage(advertiserID, 1)

	for index := 2; index <= pag.TotalPages; index++ {
		currentProduct, _ := ps.GetProductPage(advertiserID, index)
		products = append(products, currentProduct...)
		time.Sleep(ps.SleepTime)
	}

	return products
}
