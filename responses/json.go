package responses

import (
	"encoding/json"
	"net/http"
)

func SuccessJson(w http.ResponseWriter,statusCode int,data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ErrorJson(w http.ResponseWriter,statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((statusCode))
	json.NewEncoder(w).Encode(err)
}
