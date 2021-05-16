package coinservice

import (
	"fmt"

	"github.com/morscino/gigo/model/coinmodel"
	"gorm.io/gorm"
)

type CoinService struct {
	db *gorm.DB
}

type CoinRepository interface {
	StoreCoin(coin coinmodel.Coin)
	GetCoinBySymbol(symbol string, coin *coinmodel.Coin) []coinmodel.Coin
}

func NewCoinService(database *gorm.DB) CoinRepository {
	return &CoinService{db: database}
}

func (c CoinService) StoreCoin(coin coinmodel.Coin) {
	newCoindata := c.db.Create(&coin)

	if newCoindata.Error != nil {
		fmt.Print("there was an error")
	}
	//return coin
}

func (c CoinService) GetCoinBySymbol(symbol string, coin *coinmodel.Coin) []coinmodel.Coin {
	var singleCoinData []coinmodel.Coin
	c.db.Where("symbol = ?", symbol).Find(&singleCoinData)

	return singleCoinData

}
