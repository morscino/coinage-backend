package coinmodel

import (
	"time"

	"github.com/google/uuid"
)

type Coin struct {
	ID                 uuid.UUID `json:"ID" gorm:"type:uuid;primaryKey"`
	Symbol             string    `json:"Symbol" gorm:"type:varchar(100);not null"`
	PriceChange        string    `json:"PriceChange" gorm:"type:varchar(100);not null"`
	PriceChangePercent string    `json:"PriceChangePercent" gorm:"type:varchar(100);not null"`
	WeightedAvgPrice   string    `json:"WeightedAvgPrice" gorm:"type:varchar(100);not null"`
	PrevClosePrice     string    `json:"PrevClosePrice" gorm:"type:varchar(100);not null"`
	LastPrice          string    `json:"LastPrice" gorm:"type:varchar(100);not null"`
	LastQty            string    `json:"LastQty" gorm:"type:varchar(100);not null"`
	BidPrice           string    `json:"BidPrice" gorm:"type:varchar(100);not null"`
	AskPrice           string    `json:"AskPrice" gorm:"type:varchar(100);not null"`
	AskQty             string    `json:"AskQty" gorm:"type:varchar(100);not null"`
	OpenPrice          string    `json:"OpenPrice" gorm:"type:varchar(100);not null"`
	HighPrice          string    `json:"HighPrice" gorm:"type:varchar(100);not null"`
	LowPrice           string    `json:"LowPrice" gorm:"type:varchar(100);not null"`
	Volume             string    `json:"Volume" gorm:"type:varchar(100);not null"`
	QuoteVolume        string    `json:"QuoteVolume" gorm:"type:varchar(100);not null"`
	OpenTime           int       `json:"OpenTime" gorm:"type:bigint;not null"`
	CloseTime          int       `json:"CloseTime" gorm:"type:bigint;not null"`
	FirstId            int       `json:"FirstId" gorm:"type:int;not null"`
	LastId             int       `json:"LastId" gorm:"type:int;not null"`
	Count              int       `json:"Count" gorm:"type:int;not null"`
	CreatedAt          time.Time `json:"CreatedAt"`
	UpdatedAt          time.Time `json:"UpdatedAt"`
}
