package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mihirkelkar/bss-service/controllers"
	"github.com/mihirkelkar/bss-service/models"
)

func main() {
	//define the command line arguments.
	var redisConfig string
	flag.StringVar(&redisConfig, "redisconfig", "redis.json", "Provide the JSON file that has the redis config")

	config := NewConfig()
	config.ReadRedisJSON(redisConfig)

	r := mux.NewRouter()
	services, err := models.NewServices(config.ReturnConfig())
	if err != nil {
		panic(err)
	}
	defer services.Close()

	// if we can form the services correctly, then go ahead
	// and pass the service to the barcode controller.
	brCtrl := controllers.NewBarcodeController(services.ProductService)

	r.HandleFunc("/barcode", brCtrl.FetchBarcode).Methods("GET")
	r.HandleFunc("/barcode", brCtrl.AddBarcode).Methods("POST")

	http.ListenAndServe(":3000", r)

}
