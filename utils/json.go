package utils

import (
	"encoding/json"
	"net/http"
)

func WriteAsJson(w http.ResponseWriter, v interface{}) {
	jsonPayload, err := json.Marshal(v)
	CheckFatal(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonPayload)
}
