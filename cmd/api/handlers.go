package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/qlfzn/etf-go/internal/services"
)

func GetProfilesHandler (w http.ResponseWriter, r *http.Request) {
	var payload map[string]float64

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Printf("invalid json fields: %s", err)
		return
	}

	profiles, err := services.GetETFProfile(payload)
	if err != nil {
		log.Printf("error occurred: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(profiles)
}