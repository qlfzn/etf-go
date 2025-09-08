package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/qlfzn/etf-go/internal/models"
)

// client to call external API
type Client struct {
	APIKey string
}

func (c *Client) GetETFProfile(symbol string) (models.ETFProfile, error) {
	var etfProfile models.ETFProfile

	baseURL := "https://www.alphavantage.co/query"

	params := url.Values{} 
	params.Add("function", "ETF_Profile")
	params.Add("symbol", symbol)
	params.Add("apikey", c.APIKey)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// make request
	resp, err := http.Get(fullURL)
	if err != nil {
		return etfProfile, err
	}
	defer resp.Body.Close()

	// decode json
	err = json.NewDecoder(resp.Body).Decode(&etfProfile)
	if err != nil {
		return etfProfile, nil
	}

	return etfProfile, nil
}