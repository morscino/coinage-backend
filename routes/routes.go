package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/morscino/gigo/facade"
)

type CoinRoute struct {
	CoinFacade facade.CoinFacade
}

func NewCoinRoute(cf facade.CoinFacade) *CoinRoute {
	return &CoinRoute{CoinFacade: cf}
}

func (c CoinRoute) CoinRoutes(router *gin.Engine) {

	CoinsGroup := router.Group("/coins")
	{
		CoinsGroup.GET("/", c.CoinFacade.StoreCoin)
		CoinsGroup.GET("/all", c.CoinFacade.GetCoins)
		CoinsGroup.GET("/:symbol", c.CoinFacade.GetCoinBySymbol)

	}

}
