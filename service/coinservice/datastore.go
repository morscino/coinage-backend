package coinservice

import "github.com/morscino/gigo/model/coinmodel"

type CoinService struct {
	db interface{}
}

type CoinRepository interface {
	StoreCoin(coin coinmodel.Coin) coinmodel.Coin
}

func NewCoinService(database interface{}) CoinRepository {
	return &CoinService{db: database}
}

func (c CoinService) StoreCoin(coin coinmodel.Coin) coinmodel.Coin {
	return coin
}
