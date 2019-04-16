package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

type APIProduct struct {
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
	Products []APIProduct `json: products`
}

//BarcodeLookup : Interface to lookup third party barcodes and process
// them into formats compatible for our API
type BarcodeLookup interface {
	ThirdPartyAPI //This is the interface that actually connects to the API
	//and returns a response
	LookupBarcode(barcode string) (*Product, error)
	LookupDetailedBarcode(barcode string) (*LookupApiResponse, error)
}

type barcodeLookup struct {
	ThirdPartyAPI
}

func (b *barcodeLookup) loadAPIResponse(resp io.Reader) (*LookupApiResponse, error) {
	if resp == nil {
		return nil, errors.New("Error: Invalid Response Body")
	}
	result, err := ioutil.ReadAll(resp)
	if err != nil {
		return nil, errors.New("Error: Response body cannot be read")
	}
	apiresp := LookupApiResponse{}
	err = json.Unmarshal(result, &apiresp)
	if err != nil {
		return nil, errors.New("Error: Error Unmarshallling the Response")
	}
	return &apiresp, nil
}

func (b *barcodeLookup) convertToProduct(lApi *LookupApiResponse) (*Product, error) {
	var product Product
	if lApi == nil {
		return nil, errors.New("Error: The lookup api response object was nil")
	}
	product.ProductName = lApi.Products[0].ProductName
	product.Catalog = 0
	product.SearchTerm = product.ProductName
	product.Upc = lApi.Products[0].BarcodeNumber
	return &product, nil
}

//LookupBarcode : This function takes in a barcode, calls the third party Api
//interface to lookup the barcode in the third party service, convert the
//response to a product response and then returns it to the upper controller
// as a product that our service can consume.
func (b *barcodeLookup) LookupBarcode(barcode string) (*Product, error) {
	respBody, err := b.ThirdPartyAPI.queryAPI(barcode)
	if err != nil {
		return nil, err
	}
	detailPrd, err := b.loadAPIResponse(respBody)
	if err != nil {
		return nil, err
	}
	prd, err := b.convertToProduct(detailPrd)
	if err != nil {
		return nil, err
	}
	return prd, nil
}

func (b *barcodeLookup) LookupDetailedBarcode(barcode string) (*LookupApiResponse, error) {
	return &LookupApiResponse{}, nil
}

//NewBarcodeLookup : wrapper function that can return a new barcode API service
func NewBarcodeLookup(apiurl string) (BarcodeLookup, error) {
	tpa := NewThirdPartyAPI(apiurl)
	return &barcodeLookup{ThirdPartyAPI: tpa}, nil
}

//ThirdPartyAPI : This is the interface that will actually interact with the third party API
//We can mock this when this interface when unittesting.
//BarcodeLookup : interacts with the BarcodeLookup API
type ThirdPartyAPI interface {
	queryAPI(barcode string) (io.Reader, error)
}

type thirdPartyAPI struct {
	fullAPIUrl string
}

func (tp *thirdPartyAPI) queryAPI(barcode string) (io.Reader, error) {
	//make the full api url by formatting the barcode in the API
	fullBarCodeURL := tp.fullAPIUrl + fmt.Sprintf("&barcode=%s", barcode)
	req, err := http.NewRequest("GET", fullBarCodeURL, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if (resp.StatusCode == http.StatusNotFound) || (resp.StatusCode == http.StatusBadRequest) {
		return nil, errors.New("No response could be found for this product")
	}

	return resp.Body, err
}

//NewThirdPartyAPI : return an object that can fit that interface.
func NewThirdPartyAPI(apiurl string) ThirdPartyAPI {
	return &thirdPartyAPI{
		fullAPIUrl: apiurl,
	}
}
