package controllers

import (
	"net/http"

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
	data, _ := bs.ps.ByUpc("123456")
	views.Render(w, r, data)
}
