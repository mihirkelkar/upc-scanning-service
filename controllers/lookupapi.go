package controllers

import (
	"net/http"

	"github.com/mihirkelkar/bss-service/models"
)

//Stores : JSON object that gets returned from the third party API
type Stores struct {
	Storename      string `json:"store_name"`
	StorePrice     string `json:"store_price"`
	ProductUrl     string `json:"product_url"`
	CurrencyCode   string `json:"currency_code"`
	CurrencySymbol string `json:"currency_symbol"`
}

//Reviews : JSON object that gets returned from the third party API
type Reviews struct {
	Name     string `json:"name"`
	Rating   string `json:"rating"`
	Title    string `json:"title"`
	Review   string `json:"review"`
	Datetime string `json:"datetime"`
}

type Product struct {
	BarcodeNumber   string    `json:"barcode_number"`
	BarcodeType     string    `json:"barcode_type"`
	BarcodeFormat   string    `json:"barcode_formats"`
	Mpn             string    `json:"mpn"`
	Model           string    `json:"model"`
	Asin            string    `json:"asin"`
	ProductName     string    `json:"product_name"`
	Title           string    `json:"title"`
	Category        string    `json:"category"`
	Manufacturer    string    `json:"manufacturer"`
	Brand           string    `json:"brand"`
	Label           string    `json:"label"`
	Author          string    `json:"author"`
	Publisher       string    `json:"publisher"`
	Artist          string    `json:"artist"`
	Actor           string    `json:"actor"`
	Director        string    `json:"director"`
	Studio          string    `json:"studio"`
	Genre           string    `json:"genre"`
	AudienceRating  string    `json:"audience_rating"`
	Ingredients     string    `json:"ingredients"`
	NutritionFacts  string    `json:"nutrition_facts"`
	Color           string    `json:"color"`
	Format          string    `json:"format"`
	PackageQuantity string    `json:"package_quantity"`
	Size            string    `json:"size"`
	Length          string    `json:"length"`
	Width           string    `json:"width"`
	Height          string    `json:"height"`
	Weight          string    `json:"weight"`
	ReleaseDate     string    `json:"release_date"`
	Description     string    `json:"description"`
	Features        []string  `json:"features"`
	Images          []string  `json:"images"`
	Stores          []Stores  `json:"stores"`
	Reviews         []Reviews `json:"reviews"`
}

//LookupApiResponse : Struct for response of API
type LookupApiResponse struct {
	Products []Product `json: products`
}

//BarcodeLookup : Interface to lookup third party barcodes and process
// them into formats compatible for our API
type BarcodeLookup interface {
	ThirdPartyAPI //This is the interface that actually connects to the API
	//and returns a response
	LookupBarcode(barcode string) (*models.Product, error)
	LookupDetailedBarcode(barcode string) (*LookupApiResponse, error)
}

/*
type barcodeLookup struct{

}


func (b *barcodeLookup) LookupBarcode(barcode string) (*models.Product, error) {
	return &models.Product{}, nil
}

func (b *barcodeLookup) LookupDetailedBarcode(barcode string) (*LookupApiResponse, error) {
	return &LookupApiResponse{}, nil
}

func NewBarcodeLookupApi() (BarcodeLookup, error) {
	return &barcodeLookupAPI{}, nil
}
*/

//This is the interface that will actually interact with the third party API
//We can mock this when this interface when unittesting.
//BarcodeLookup : interacts with the BarcodeLookup API
type ThirdPartyAPI interface {
	queryAPI(barcode string) (*http.Response, error)
}

type thirdPartyAPI struct {
	apikey string
	apiurl string
}

func (tp *thirdPartyAPI) queryAPI(barcode string) (*http.Response, error) {
	return nil, nil
}
