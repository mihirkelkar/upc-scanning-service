package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mihirkelkar/bss-service/models"
)

var brctrl *BarcodeController

func TestFetchBarcode(t *testing.T) {
	//init a response recorder
	rr := httptest.NewRecorder()

	//declaree an endpoint.
	r, err := http.NewRequest("GET", "/barcode", nil)
	if err != nil {
		t.Fatal(err)
	}
	brctrl.FetchBarcode(rr, r)
	rs := rr.Result()

	//check status code. A call to just the barcode API is a bad request.
	if rs.StatusCode != http.StatusBadRequest {
		t.Errorf("Error: Expected %d recieved %d", http.StatusOK, rs.StatusCode)
	}

	//now check the same request with an empty upc parameter.
	r, err = http.NewRequest("GET", "/barcode?upc", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	brctrl.FetchBarcode(rr, r)
	rs = rr.Result()
	if rs.StatusCode != http.StatusBadRequest {
		t.Errorf("Error: Expected %d recieved %d", http.StatusBadRequest, rs.StatusCode)
	}

	//now check an actual correct request.
	r, err = http.NewRequest("GET", "/barcode?upc=123456", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	brctrl.FetchBarcode(rr, r)
	rs = rr.Result()
	if rs.StatusCode != http.StatusOK {
		t.Errorf("Error: Expected %d recieved %d", http.StatusOK, rs.StatusCode)
	}

	//now lets check if we pass in a non alphanumeric string as a upc paramter
	rr = httptest.NewRecorder()
	r, err = http.NewRequest("GET", "/barcode?upc=abcdef", nil)
	if err != nil {
		t.Fatal(err)
	}
	brctrl.FetchBarcode(rr, r)
	rs = rr.Result()
	if rs.StatusCode != http.StatusBadRequest {
		t.Errorf("Error: Expected %d recieved %d", http.StatusBadRequest, rs.StatusCode)
	}
}

func TestMain(m *testing.M) {
	//setup the testing here.
	var configmap = make(map[string]string)
	configmap["address"] = "localhost:6379"
	configmap["password"] = "****" //this is not the real password duh
	configmap["database"] = "0"
	configmap["apiurl"] = "https://api.barcodelookup.com/v2/products?formatted=y&key=%s"
	configmap["apikey"] = "****"

	services, err := models.NewServices(configmap)
	if err != nil {
		os.Exit(1)
	}
	defer services.Close()
	brctrl = NewBarcodeController(services.ProductService)
	os.Exit(m.Run())
}
