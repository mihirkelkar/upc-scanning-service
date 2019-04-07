package controllers

import (
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
		http.NotFound(w, r)
	}
	views.Render(w, r, data)
}
