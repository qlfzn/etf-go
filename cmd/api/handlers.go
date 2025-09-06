package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type CompareETFPayload struct {
	Investments map[string]float64 `json:"investments"`
}

func CompareETFHandler(w http.ResponseWriter, r *http.Request) {
	var payload CompareETFPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("invalid json fields: %s", err)
	}

	// TODO: call service here

}