package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mihirkelkar/bss-service/controllers"
	"github.com/mihirkelkar/bss-service/models"
)

func SetupTest() *mux.Router {
	r := mux.NewRouter()
	services, err := models.NewServices()
	if err != nil {
		panic(err)
	}
	defer services.Close()
	brCtrl := controllers.NewBarcodeController(services.ProductService)
	r.HandleFunc("/barcode", brCtrl.FetchBarcode).Methods("GET")
	return r
}

func TestBarcodeEndPointInvalid(t *testing.T) {
	req, _ := http.NewRequest("GET", "/barcode", nil)
	response := httptest.NewRecorder()
	SetupTest().ServeHTTP(response, req)
	if response.Code != http.StatusBadRequest {
		fmt.Errorf("Failure: Expected %d recvieved %d", http.StatusBadRequest, response.Code)
	}
}

func TestBarcodeEndPointInvalidParams(t *testing.T) {
	req, _ := http.NewRequest("GET", "/barcode?upc=", nil)
	response := httptest.NewRecorder()
	SetupTest().ServeHTTP(response, req)
	if response.Code != http.StatusBadRequest {
		fmt.Errorf("Failure: Expected %d recvieved %d", http.StatusBadRequest, response.Code)
	}
}

func TestBarcodeEndPointValid(t *testing.T) {
	req, _ := http.NewRequest("GET", "/barcode?upc=123456", nil)
	response := httptest.NewRecorder()
	SetupTest().ServeHTTP(response, req)
	if response.Code != http.StatusFound {
		fmt.Errorf("Failure: Expected %d recvieved %d", http.StatusFound, response.Code)
	}
}
