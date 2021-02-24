package apisdk

import (
	"fmt"
)

// FormatTickerResults gets an APIResponse
// and returns formatted string that is ready to be sent as a bots response
func FormatTickerResults(tickerInfo APIResponse) string {
	match := tickerInfo.QuoteSummary.Result[0]

	if match.Price.RegularMarketPrice.Fmt == "" {
		return " ¯\\_(ツ)_/¯ "
	}

	graphEmoji := "📈"

	if match.Price.RegularMarketChangePercent.Raw < 0 {
		graphEmoji = "📉"
	}

	result := fmt.Sprintf(
		"$%s __%s__\n%s *%s%s* (%s)\n🔼%s%s 🔽%s%s",
		match.Price.Symbol,
		match.Price.ShortName,
		graphEmoji,
		match.Price.CurrencySymbol,
		match.Price.RegularMarketPrice.Fmt,
		match.Price.RegularMarketChangePercent.Fmt,
		match.Price.CurrencySymbol,
		match.Price.RegularMarketDayHigh.Fmt,
		match.Price.CurrencySymbol,
		match.Price.RegularMarketDayLow.Fmt,
	)

	return result
}
