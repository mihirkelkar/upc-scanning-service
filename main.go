package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihirkelkar/bss-service/controllers"
	"github.com/mihirkelkar/bss-service/models"
)

func main() {

	r := mux.NewRouter()
	services, err := models.NewServices()
	if err != nil {
		panic(err)
	}

	// if we can form the services correctly, then go ahead
	// and pass the service to the barcode controller.
	brCtrl := controllers.NewBarcodeController(services.ProductService)

	r.HandleFunc("/barcode", brCtrl.FetchBarcode).Methods("GET")

	http.ListenAndServe(":3000", r)

}
