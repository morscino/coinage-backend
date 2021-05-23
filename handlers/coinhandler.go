package handlers

import (
	"fmt"

	"time"

	"github.com/google/uuid"
	"github.com/imroc/req"
	"github.com/morscino/gigo/model/coinmodel"
	"github.com/morscino/gigo/service/coinservice"
	"github.com/morscino/gigo/utility/log"
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
	CoinService coinservice.CoinRepository
}

func NewCoinHandler(c coinservice.CoinRepository) CoinHandler {
	return CoinHandler{CoinService: c}
}

func (c CoinHandler) RunCron(d time.Duration, t int) {

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
	allCoins := c.GetCoins()

	for _, singleCoin := range allCoins {
		id := uuid.New()
		coin := coinmodel.Coin{
			ID:                 id,
			Symbol:             singleCoin.Symbol,
			PriceChange:        singleCoin.PriceChange,
			PriceChangePercent: singleCoin.PriceChangePercent,
			WeightedAvgPrice:   singleCoin.WeightedAvgPrice,
			PrevClosePrice:     singleCoin.PrevClosePrice,
			LastPrice:          singleCoin.LastPrice,
			LastQty:            singleCoin.LastQty,
			BidPrice:           singleCoin.BidPrice,
			AskPrice:           singleCoin.AskPrice,
			AskQty:             singleCoin.AskQty,
			OpenPrice:          singleCoin.OpenPrice,
			HighPrice:          singleCoin.HighPrice,
			LowPrice:           singleCoin.LowPrice,
			Volume:             singleCoin.Volume,
			QuoteVolume:        singleCoin.QuoteVolume,
			OpenTime:           singleCoin.OpenTime,
			CloseTime:          singleCoin.CloseTime,
			FirstId:            singleCoin.FirstId,
			LastId:             singleCoin.LastId,
			Count:              singleCoin.Count,
			CreatedAt:          time.Now(),
		}

		c.CoinService.StoreCoin(coin)

	}
	fmt.Printf("Coins stored after %v seconds \n", t)
	//return allCoins
}

func (c CoinHandler) GetCoins() []CoinResponse {

	//curl to get coins
	url := "https://api.binance.com/api/v3/ticker/24hr"

	resp, err := req.Get(url)
	if err != nil {
		log.Error("Error fetching data: %v", err.Error())
		//fmt.Println(err)

	}

	if err := resp.ToJSON(&response); err != nil {
		//fmt.Println(err)
		log.Error("Error binding data: %v", err.Error())

	}

	return response
}

func (c CoinHandler) GetCoinBySymbol(symbol string) *[]coinmodel.Coin {
	var coin *coinmodel.Coin
	result := c.CoinService.GetCoinBySymbol(symbol, coin)

	return &result
}
