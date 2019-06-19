// Package response contains functions related to how api end point generate
// response.
package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WithJSONStr generates JSON response with string.
func WithJSONStr(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

// WithJSONBytes generates JSON response with bytes.
func WithJSONBytes(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

// WithJSONData generates JSON response with data.
func WithJSONData(w http.ResponseWriter, data interface{}, code int) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{error: %q}", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(body)
}
