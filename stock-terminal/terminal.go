package stockterminal

import (
	"log"
	"os"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var client *sdk.StreamingClient

// InitStockClient initializes Tinkoff Invest OpenAPI client
func InitStockClient(token string) {
	logger := log.New(os.Stdout, "[invest-openapi-go-sdk]", log.LstdFlags)

	client, err := sdk.NewStreamingClient(logger, token)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
}
