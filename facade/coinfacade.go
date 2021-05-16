package facade

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/morscino/gigo/handlers"
)

type CoinFacade struct {
	ctx         context.Context
	CoinHandler handlers.CoinHandler
}

func NewCoinFacade(ctx context.Context, ch handlers.CoinHandler) *CoinFacade {
	return &CoinFacade{
		ctx:         ctx,
		CoinHandler: ch,
	}
}

func (coin CoinFacade) StoreCoin(c *gin.Context) {

	coin.CoinHandler.RunCron(30000*time.Millisecond, 20000/1000)
	c.JSON(http.StatusOK, gin.H{"data": "coin stored"})
}

func (coin CoinFacade) GetCoins(c *gin.Context) {

	allCoins := coin.CoinHandler.GetCoins()
	c.JSON(http.StatusOK, gin.H{"data": allCoins})
}

func (coin CoinFacade) GetCoinBySymbol(c *gin.Context) {
	symbol := strings.ToUpper(c.Param("symbol"))

	singleCoinData := coin.CoinHandler.GetCoinBySymbol(symbol)

	c.JSON(http.StatusOK, gin.H{"data": singleCoinData})

}
