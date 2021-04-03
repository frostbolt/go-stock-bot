package formatters

import (
	"fmt"

	"github.com/frostbolt/go-stock-bot/apisdk"
)

// FormatCurrency gets currency code and the value and formats it properly
func FormatCurrency(currencyCode string, amount float64) string {
	symbol, isDefined := CurrencySymbols[currencyCode]

	if !isDefined {
		symbol = currencyCode
	}

	sign := ""
	if amount < 0 {
		sign = "-"
	}

	return fmt.Sprintf("%s%s%0.2f", sign, symbol, amount)
}

// FormatTickerResults gets an APIResponse
// and returns formatted string that is ready to be sent as a bots response
func FormatTickerResults(tickerInfo apisdk.QuoteAPIResponse) string {
	match := tickerInfo.QuoteSummary.Result[0]

	if match.Price.RegularMarketPrice.Fmt == "" {
		return " ¯\\_(ツ)_/¯ "
	}

	graphEmoji := "📈"

	if match.Price.RegularMarketChangePercent.Raw < 0 {
		graphEmoji = "📉"
	}

	var currency string = match.Price.Currency

	result := fmt.Sprintf(
		"$%s __%s__\n%s*%s* (%s)\n`Today:` 🔼%s 🔽%s\n`52 wk:` 🔼%s 🔽%s",
		match.Price.Symbol,
		match.Price.ShortName,
		graphEmoji,
		FormatCurrency(currency, match.Price.RegularMarketPrice.Raw),
		match.Price.RegularMarketChangePercent.Fmt,
		FormatCurrency(currency, match.Price.RegularMarketDayHigh.Raw),
		FormatCurrency(currency, match.Price.RegularMarketDayLow.Raw),
		FormatCurrency(currency, match.SummaryDetail.FiftyTwoWeekHigh.Raw),
		FormatCurrency(currency, match.SummaryDetail.FiftyTwoWeekLow.Raw),
	)

	return result
}
