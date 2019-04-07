package views

import (
	"encoding/json"
	"net/http"
)

//Render : Renders the data provided.
//Note that it is not the render function's responsibility on making a decision
//on whether to render data or not or to analyze the data at all.
//it is the render function's repsonsibility to render data.
//all calls about empty data . malformed data should be made from the controller.
func Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}
