package views

import (
	"encoding/json"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
