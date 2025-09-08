package services

import (
	"log"

	"github.com/qlfzn/etf-go/internal/models"
)

// orchestrate fetching + calculation
func CompareETF(investments map[string]float64) {
	var profiles []models.ETFProfile
	totalAmountInvested := 0.0

	client := &Client{
		APIKey: "",
	}

	for symbol, amount := range investments {
		profile, err := client.GetETFProfile(symbol)
		if err != nil {
			log.Printf("error fetching profile: %s", err)
		}

		profile.InvestedAmount = amount

		profiles = append(profiles, profile)
		totalAmountInvested += amount

	}

}