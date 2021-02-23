package apisdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/frostbolt/go-stock-bot/cache"
	log "github.com/sirupsen/logrus"
)

var client = &http.Client{Timeout: 10 * time.Second}

func makeAnAPICall(url string, target interface{}) error {
	body, err := cache.CachedFunctionCall(url, func() (string, error) {
		resp, err := client.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		var bodyString string

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString = string(bodyBytes)
		}

		return bodyString, nil
	})

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(body), target)
}

// GetInfoOnTicker calls finance api and returns a decoded response object
func GetInfoOnTicker(ticker string) (APIResponse, error) {
	requestURL := fmt.Sprintf("https://query1.finance.yahoo.com/v11/finance/quoteSummary/%s?modules=price", ticker)
	result := new(APIResponse)
	err := makeAnAPICall(requestURL, result)

	if result.QuoteSummary.Result == nil {
		return *result, errors.New("Invalid request")
	}

	return *result, err
}
