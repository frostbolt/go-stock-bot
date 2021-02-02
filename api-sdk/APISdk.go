package apisdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	resp, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// GetInfoOnTicker calls finance api and returns a decoded response object
func GetInfoOnTicker(ticker string) (APIResponse, error) {
	requestURL := fmt.Sprintf("https://query1.finance.yahoo.com/v11/finance/quoteSummary/%s?modules=price", ticker)
	result := new(APIResponse)

	err := getJson(requestURL, result)

	if result.QuoteSummary.Result == nil {
		return *result, errors.New("Invalid request")
	}

	return *result, err
}
