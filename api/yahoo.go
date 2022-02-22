package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

type QuoteStructure struct {
	Symbol                     string
	ShortName                  string
	RegularMarketPrice         float64 // WE EXPECT HIGH NUMBERS, RIGHT
	RegularMarketChange        float64
	RegularMarketChangePercent float64
	RegularMarketVolume        int
	PreMarketPrice             float64
	PreMarketChange            float64
	PreMarketChangePercent     float64
	PostMarketPrice            float64
	PostMarketChange           float64
	PostMarketChangePercent    float64
	MarketState                string
	Currency                   string
	Exchange                   string
}

type Result struct {
	Result []QuoteStructure `json:"result"`
}

type QuoteResponse struct {
	QuoteResponse Result `json:"quoteResponse"`
}

func Quote(symbol string) QuoteStructure {
	api := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/quote?symbols=%s", symbol)
	resp, err := http.Get(api)

	if err != nil {
		log.Fatal(err)
	}

	respData, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var quote QuoteResponse

	jsonErr := json.Unmarshal(respData, &quote)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return quote.QuoteResponse.Result[0]
}

func (stock *QuoteStructure) Price() float64 {
	switch stock.MarketState {
	case "PRE":
		return stock.PreMarketPrice
	case "POST":
		return stock.PostMarketPrice
	default:
		return stock.RegularMarketPrice
	}
}

func (stock *QuoteStructure) PriceState() bool {
	switch stock.MarketState {
	case "PRE":
		return !math.Signbit(stock.PreMarketChange)
	case "POST":
		return !math.Signbit(stock.PostMarketChange)
	default:
		return !math.Signbit(stock.RegularMarketChange)
	}
}

func (stock *QuoteStructure) State() string {
	switch stock.MarketState {
	case "PRE":
		return "Pre"
	case "POST":
		return "After hours"
	case "CLOSED":
		return "Closed"
	default:
		return "Open"
	}
}

func (stock *QuoteStructure) MarketChangePercent() string {
	switch stock.MarketState {
	case "PRE":
		return fmt.Sprintf("%.2f", stock.PreMarketChangePercent) + "%"
	case "POST":
		return fmt.Sprintf("%.2f", stock.PostMarketChangePercent) + "%"
	default:
		return fmt.Sprintf("%.2f", stock.RegularMarketChangePercent) + "%"
	}
}

func (stock *QuoteStructure) MarketChange() string {
	switch stock.MarketState {
	case "PRE":
		return fmt.Sprintf("%.2f ", stock.PreMarketChange) + stock.Currency
	case "POST":
		return fmt.Sprintf("%.2f ", stock.PostMarketChange) + stock.Currency
	default:
		return fmt.Sprintf("%.2f ", stock.RegularMarketChange) + stock.Currency
	}
}

func (stock *QuoteStructure) MarketVolume() int {
	if stock.MarketState == "REGULAR" {
		return stock.RegularMarketVolume
	}

	return 0
}

func (stock *QuoteStructure) Name() string {
	return stock.ShortName
}

func (stock *QuoteStructure) StockCurrency() string {
	return stock.Currency
}
