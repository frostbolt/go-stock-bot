package apisdk

// QuoteAPIResponse describes the respose from yahoo finance api `quoteSummary` method
type QuoteAPIResponse struct {
	QuoteSummary struct {
		Result []struct {
			SummaryDetail struct {
				MaxAge    int `json:"maxAge"`
				PriceHint struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"priceHint"`
				PreviousClose struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"previousClose"`
				Open struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"open"`
				DayLow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"dayLow"`
				DayHigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"dayHigh"`
				RegularMarketPreviousClose struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPreviousClose"`
				RegularMarketOpen struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketOpen"`
				RegularMarketDayLow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayLow"`
				RegularMarketDayHigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayHigh"`
				DividendRate struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"dividendRate"`
				DividendYield struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"dividendYield"`
				ExDividendDate struct {
					Raw int    `json:"raw"`
					Fmt string `json:"fmt"`
				} `json:"exDividendDate"`
				PayoutRatio struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"payoutRatio"`
				FiveYearAvgDividendYield struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"fiveYearAvgDividendYield"`
				Beta struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"beta"`
				TrailingPE struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"trailingPE"`
				ForwardPE struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"forwardPE"`
				Volume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"volume"`
				RegularMarketVolume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"regularMarketVolume"`
				AverageVolume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"averageVolume"`
				AverageVolume10Days struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"averageVolume10days"`
				AverageDailyVolume10Day struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"averageDailyVolume10Day"`
				Bid struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"bid"`
				Ask struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"ask"`
				BidSize struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"bidSize"`
				AskSize struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"askSize"`
				MarketCap struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"marketCap"`
				Yield struct {
				} `json:"yield"`
				YtdReturn struct {
				} `json:"ytdReturn"`
				TotalAssets struct {
				} `json:"totalAssets"`
				ExpireDate struct {
				} `json:"expireDate"`
				StrikePrice struct {
				} `json:"strikePrice"`
				OpenInterest struct {
				} `json:"openInterest"`
				FiftyTwoWeekLow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"fiftyTwoWeekLow"`
				FiftyTwoWeekHigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"fiftyTwoWeekHigh"`
				PriceToSalesTrailing12Months struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"priceToSalesTrailing12Months"`
				FiftyDayAverage struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"fiftyDayAverage"`
				TwoHundredDayAverage struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"twoHundredDayAverage"`
				TrailingAnnualDividendRate struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"trailingAnnualDividendRate"`
				TrailingAnnualDividendYield struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"trailingAnnualDividendYield"`
				NavPrice struct {
				} `json:"navPrice"`
				Currency     string      `json:"currency"`
				FromCurrency interface{} `json:"fromCurrency"`
				ToCurrency   interface{} `json:"toCurrency"`
				LastMarket   interface{} `json:"lastMarket"`
				Volume24Hr   struct {
				} `json:"volume24Hr"`
				VolumeAllCurrencies struct {
				} `json:"volumeAllCurrencies"`
				CirculatingSupply struct {
				} `json:"circulatingSupply"`
				Algorithm interface{} `json:"algorithm"`
				MaxSupply struct {
				} `json:"maxSupply"`
				StartDate struct {
				} `json:"startDate"`
				Tradeable bool `json:"tradeable"`
			} `json:"summaryDetail"`
			Price struct {
				MaxAge                 int `json:"maxAge"`
				PreMarketChangePercent struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"preMarketChangePercent"`
				PreMarketChange struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"preMarketChange"`
				PreMarketTime  int `json:"preMarketTime"`
				PreMarketPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"preMarketPrice"`
				PreMarketSource         string `json:"preMarketSource"`
				PostMarketChangePercent struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"postMarketChangePercent"`
				PostMarketChange struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"postMarketChange"`
				PostMarketTime  int `json:"postMarketTime"`
				PostMarketPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"postMarketPrice"`
				PostMarketSource           string `json:"postMarketSource"`
				RegularMarketChangePercent struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketChangePercent"`
				RegularMarketChange struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketChange"`
				RegularMarketTime int `json:"regularMarketTime"`
				PriceHint         struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"priceHint"`
				RegularMarketPrice struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPrice"`
				RegularMarketDayHigh struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayHigh"`
				RegularMarketDayLow struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketDayLow"`
				RegularMarketVolume struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"regularMarketVolume"`
				AverageDailyVolume10Day struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"averageDailyVolume10Day"`
				AverageDailyVolume3Month struct {
					Raw     int    `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"averageDailyVolume3Month"`
				RegularMarketPreviousClose struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketPreviousClose"`
				RegularMarketSource string `json:"regularMarketSource"`
				RegularMarketOpen   struct {
					Raw float64 `json:"raw"`
					Fmt string  `json:"fmt"`
				} `json:"regularMarketOpen"`
				StrikePrice struct {
				} `json:"strikePrice"`
				OpenInterest struct {
				} `json:"openInterest"`
				Exchange              string      `json:"exchange"`
				ExchangeName          string      `json:"exchangeName"`
				ExchangeDataDelayedBy int         `json:"exchangeDataDelayedBy"`
				MarketState           string      `json:"marketState"`
				QuoteType             string      `json:"quoteType"`
				Symbol                string      `json:"symbol"`
				UnderlyingSymbol      interface{} `json:"underlyingSymbol"`
				ShortName             string      `json:"shortName"`
				LongName              string      `json:"longName"`
				Currency              string      `json:"currency"`
				QuoteSourceName       string      `json:"quoteSourceName"`
				CurrencySymbol        string      `json:"currencySymbol"`
				FromCurrency          interface{} `json:"fromCurrency"`
				ToCurrency            interface{} `json:"toCurrency"`
				LastMarket            interface{} `json:"lastMarket"`
				Volume24Hr            struct {
				} `json:"volume24Hr"`
				VolumeAllCurrencies struct {
				} `json:"volumeAllCurrencies"`
				CirculatingSupply struct {
				} `json:"circulatingSupply"`
				MarketCap struct {
					Raw     int64  `json:"raw"`
					Fmt     string `json:"fmt"`
					LongFmt string `json:"longFmt"`
				} `json:"marketCap"`
			} `json:"price"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteSummary"`
}

// SearchAPIResponse describes the response from yahoo finance api `search` method
type SearchAPIResponse struct {
	Explains []interface{} `json:"explains"`
	Count    int           `json:"count"`
	Quotes   []struct {
		Exchange       string  `json:"exchange,omitempty"`
		Shortname      string  `json:"shortname,omitempty"`
		QuoteType      string  `json:"quoteType,omitempty"`
		Symbol         string  `json:"symbol,omitempty"`
		Index          string  `json:"index"`
		Score          float64 `json:"score,omitempty"`
		TypeDisp       string  `json:"typeDisp,omitempty"`
		IsYahooFinance bool    `json:"isYahooFinance"`
		Longname       string  `json:"longname,omitempty"`
		Name           string  `json:"name,omitempty"`
		Permalink      string  `json:"permalink,omitempty"`
	} `json:"quotes"`
	News []struct {
		UUID                string `json:"uuid"`
		Title               string `json:"title"`
		Publisher           string `json:"publisher"`
		Link                string `json:"link"`
		ProviderPublishTime int    `json:"providerPublishTime"`
		Type                string `json:"type"`
	} `json:"news"`
	Nav []struct {
		NavName string `json:"navName"`
		NavURL  string `json:"navUrl"`
	} `json:"nav"`
	Lists []struct {
		Slug          string  `json:"slug,omitempty"`
		Name          string  `json:"name,omitempty"`
		Index         string  `json:"index"`
		Score         float64 `json:"score"`
		Type          string  `json:"type"`
		BrandSlug     string  `json:"brandSlug,omitempty"`
		PfID          string  `json:"pfId,omitempty"`
		ID            string  `json:"id,omitempty"`
		Title         string  `json:"title,omitempty"`
		CanonicalName string  `json:"canonicalName,omitempty"`
	} `json:"lists"`
	ResearchReports                []interface{} `json:"researchReports"`
	TotalTime                      int           `json:"totalTime"`
	TimeTakenForQuotes             int           `json:"timeTakenForQuotes"`
	TimeTakenForNews               int           `json:"timeTakenForNews"`
	TimeTakenForAlgowatchlist      int           `json:"timeTakenForAlgowatchlist"`
	TimeTakenForPredefinedScreener int           `json:"timeTakenForPredefinedScreener"`
	TimeTakenForCrunchbase         int           `json:"timeTakenForCrunchbase"`
	TimeTakenForNav                int           `json:"timeTakenForNav"`
	TimeTakenForResearchReports    int           `json:"timeTakenForResearchReports"`
}
