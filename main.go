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
	var apiKey string
	flag.StringVar(&redisConfig, "redisconfig", "redis.json", "Provide the JSON file that has the redis config")
	flag.StringVar(&apiKey, "apikey", "tre", "Provide the third party API key for barcodeapi.com")
	flag.Parse()

	config := NewConfig()
	config.ReadRedisJSON(redisConfig)
	config.SetApiKey(apiKey)

	r := mux.NewRouter()
	services, err := models.NewServices(config.ReturnConfig(), config.GetApiKey())
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
