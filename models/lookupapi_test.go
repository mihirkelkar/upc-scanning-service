package models

import (
	"errors"
	"io"
	"os"
	"testing"
)

type mockThirdPartyAPI struct{}

func (mct *mockThirdPartyAPI) queryAPI(filename string) (io.Reader, error) {
	//This reads the file and then returns a "mock" api response.
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("No response could be found for this product")
	}
	return f, nil
}

func TestQueryAPI(t *testing.T) {
	barcode := "test/testvalidapi.json"
	checkQueryAPINoErr := func(t *testing.T, tp ThirdPartyAPI, barcode string) {
		t.Helper()
		_, err := tp.queryAPI(barcode)
		if err != nil {
			t.Errorf("Expected no error and recieved error")
		}
	}

	checkQueryAPIErr := func(t *testing.T, tp ThirdPartyAPI, barcode string) {
		t.Helper()
		_, err := tp.queryAPI(barcode)
		if err == nil {
			t.Errorf("Expected error and recieved no error")
		}
	}

	t.Run("valid file check", func(t *testing.T) {
		mc := &mockThirdPartyAPI{}
		checkQueryAPINoErr(t, mc, barcode)
		checkQueryAPIErr(t, mc, "")
	})
}

func TestLoadApiResponse(t *testing.T) {
	blp := barcodeLookup{}

	//Test 1 loadAPIResponse check for no errors first.
	f, err := os.Open("test/testvalidapi.json")
	if err != nil {
		t.Errorf("There was an error in the way this test was setup")
	}

	apiresp, err := blp.loadAPIResponse(f)
	if err != nil {
		t.Error("Error Test 1:Expected no errors and",
			" recived an error in loadAPIResponse test")
	}

	if len(apiresp.Products) > 1 {
		t.Errorf("Error Test 1: Expected 1 product and recieved %d products",
			len(apiresp.Products))
	}

	prd := apiresp.Products[0]

	if prd.BarcodeNumber != "767719012051" {
		t.Errorf("Error Test 1: Expected 767719012051 and recieved %s",
			prd.BarcodeNumber)
	}

	//Test 2 loadAPIResponse check whether an invalid api response would work.
	f, err = os.Open("test/test_invalidapi.json")
	if err != nil {
		t.Error("Error: The test setup file is not present")
	}
	apiresp, err = blp.loadAPIResponse(f)
	if err == nil {
		t.Error("Error Test 2: Expected error recieved nil")
	}

	//Test 3 loadAPIResponse check whether a completely missing input to this
	//function would work.
	apiresp, err = blp.loadAPIResponse(nil)
	if err == nil {
		t.Error("Error Test 3: Expepcted error recieved nil")
	}
}

func TestConvertToProduct(t *testing.T) {
	//Test 4. Check the conver to product function.
	checkConvertToProductVaild := func(t *testing.T, blp barcodeLookup, filename string) {
		t.Helper()
		var lapi *LookupApiResponse
		f, err := os.Open(filename)
		lapi, err = blp.loadAPIResponse(f)
		if err != nil {
			t.Error("Error Test 4: Setup failure")
		}

		prd, err := blp.convertToProduct(lapi)
		if err != nil {
			t.Errorf("Error Test 4: Expected no error but recieved %s", err.Error())
		}

		if prd.ProductName != "Beal Flyer II 10.2Mmx70M Dry Cover Rope" {
			t.Errorf("Error Test 4 : Expected product name recieved %s", prd.ProductName)
		}

		if prd.SearchTerm != "Beal Flyer II 10.2Mmx70M Dry Cover Rope" {
			t.Errorf("Error Test 4: Expected search term recieved %s", prd.SearchTerm)
		}

		if prd.Upc != "767719012051" {
			t.Errorf("Error Test 4: Expected barcode recieved %s", prd.Upc)
		}

		if prd.Catalog != 0 {
			t.Errorf("Error Test 4: Expected false in catalog")
		}
	}
	checkConvertToProductInvalid := func(t *testing.T, blp barcodeLookup, filename string) {
		t.Helper()
		prd, _ := blp.convertToProduct(nil)
		if prd != nil {
			t.Error("Error Test 4: Expected nil")
		}
	}
	blp := barcodeLookup{}
	checkConvertToProductVaild(t, blp, "test/testvalidapi.json")
	checkConvertToProductInvalid(t, blp, "test/test_invalidapi.json")
}

func TestMain(m *testing.M) {
	//setup the testing here.
	var configmap = make(map[string]string)
	configmap["address"] = "localhost:6379"
	configmap["password"] = "****" //this is not the real password duh
	configmap["database"] = "0"
	configmap["apiurl"] = "https://api.barcodelookup.com/v2/products?formatted=y&key=%s"
	configmap["apikey"] = "****"

	services, err := NewServices(configmap)
	if err != nil {
		os.Exit(1)
	}
	defer services.Close()
	os.Exit(m.Run())
}
