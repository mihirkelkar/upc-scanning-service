package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihirkelkar/bss-service/controllers"
	"github.com/mihirkelkar/bss-service/models"
)

func main() {
	//all of this should be separated out into a new file called app.go
	r := mux.NewRouter()
	services, err := models.NewServices()
	if err != nil {
		panic(err)
	}
	defer services.Close()

	// if we can form the services correctly, then go ahead
	// and pass the service to the barcode controller.
	brCtrl := controllers.NewBarcodeController(services.ProductService)

	r.HandleFunc("/barcode", brCtrl.FetchBarcode).Methods("GET")

	http.ListenAndServe(":3000", r)

}
