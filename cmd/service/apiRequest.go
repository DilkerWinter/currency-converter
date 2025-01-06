package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ExchangeRateResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func GetExchangeRateFromApi() (map[string]float64, error) {
	const apiURL = "https://openexchangerates.org/api/latest.json"
	const appID = "Your API Key"

	url := fmt.Sprintf("%s?app_id=%s", apiURL, appID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var exchangeRate ExchangeRateResponse
	err = json.Unmarshal(body, &exchangeRate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return exchangeRate.Rates, nil
}
