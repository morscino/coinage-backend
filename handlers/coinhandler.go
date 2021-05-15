package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/imroc/req"
	"github.com/morscino/gigo/service/coinservice"
)

var response []CoinResponse

type CoinResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange" `
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice" `
	LastPrice          string `json:"lastPrice" `
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice" `
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice" `
	LowPrice           string `json:"lowPrice" `
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume" `
	OpenTime           int    `json:"openTime" `
	CloseTime          int    `json:"closeTime" `
	FirstId            int    `json:"firstId" `
	LastId             int    `json:"lastId"`
	Count              int    `json:"count"`
}

type CoinHandler struct {
	CoinRepository coinservice.CoinRepository
}

func NewCoinHandler(c coinservice.CoinRepository) CoinHandler {
	return CoinHandler{CoinRepository: c}
}

func (c CoinHandler) RunCron(d time.Duration, t int) {
	//now := time.Time()
	c.storeCoin(t)
	//Run every 30 seconds
	for x := range time.Tick(d) {
		c.caller(x, t)
	}

}

func (c CoinHandler) caller(x time.Time, t int) {
	c.storeCoin(t)
}

func (c CoinHandler) storeCoin(t int) {
	//get all coins
	//allCoins := c.GetCoins()
	fmt.Printf("Coins stored after %v seconds \n", t)
	//return allCoins
}

func (c CoinHandler) GetCoins() []CoinResponse {

	//curl to get coins
	url := "https://api.binance.com/api/v3/ticker/24hr"

	resp, err := req.Get(url)
	if err != nil {
		log.Fatal(err)
		//return 0, err
	}

	if err := resp.ToJSON(&response); err != nil {
		sentry.CaptureException(err)
		//return 0, err
	}

	return response
}
