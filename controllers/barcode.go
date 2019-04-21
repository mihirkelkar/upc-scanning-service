package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mihirkelkar/bss-service/models"
	"github.com/mihirkelkar/bss-service/views"
)

type BarcodeController struct {
	ps models.ProductService
}

//NewBarcodeController : returns a new barcode controller.
func NewBarcodeController(ps models.ProductService) *BarcodeController {
	return &BarcodeController{ps: ps}
}

func (bs *BarcodeController) FetchBarcode(w http.ResponseWriter, r *http.Request) {
	//id: lets get the id out of the query.
	upc := r.URL.Query().Get("upc")
	//malformed request
	if upc == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	if _, err := strconv.Atoi(upc); err != nil {
		http.Error(w, http.StatusText(400), 400)
	}

	data, _ := bs.ps.ByUpc(upc)
	//if data is empty, then render a 404 on the API
	if data == nil {
		data, _ = bs.ps.LookupBarcode(upc)
		if data != nil {
			err := bs.ps.AddUpc(data)
			if err != nil {
				fmt.Println("Error : Error storing the product in Redis")
			}
		}
	}
	//Render a 404 if the data is empty
	if data == nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	views.Render(w, r, data)
}

func (bs *BarcodeController) AddBarcode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product models.Product

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	err = json.Unmarshal(body, &product)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
	err = bs.ps.AddUpc(&product)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	w.WriteHeader(200)
}
