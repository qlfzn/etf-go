package services

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/qlfzn/etf-go/internal/models"
)

// Fetch ETF profiles based on submitted investments
func GetETFProfile (investments map[string]float64) ([]models.ETFProfile, error){
	var profiles []models.ETFProfile
	totalAmountInvested := 0.0

	client := &Client{
		APIKey: os.Getenv("AV_API_KEY"),
	}

	for symbol, amount := range investments {
		profile, err := client.GetETFProfile(symbol)
		if err != nil {
			return nil, err
		}

		profile.AmountInvested = amount

		profiles = append(profiles, profile)
		totalAmountInvested += amount
	}

	return profiles, nil
}