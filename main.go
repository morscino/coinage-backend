package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"github.com/morscino/gigo/database/postgres"
	"github.com/morscino/gigo/facade"
	"github.com/morscino/gigo/handlers"
	"github.com/morscino/gigo/model/coinmodel"
	"github.com/morscino/gigo/routes"
	"github.com/morscino/gigo/service/coinservice"
	"github.com/morscino/gigo/utility/config"
	"github.com/morscino/gigo/utility/log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Error("Failed to load env data: %v", err)
	}
	var config config.Config
	if err := envconfig.Process("", &config); err != nil {
		log.Error("Could not load configuration data: %v", err)
	}

	db := postgres.DbConnect(config.DB)

	//migrations
	var CoinModel coinmodel.Coin
	db.AutoMigrate(&CoinModel)

	server := gin.Default()
	var ctx context.Context

	CoinRepo := coinservice.NewCoinService(db)
	CoinHandler := handlers.NewCoinHandler(CoinRepo)
	CoinFacade := *facade.NewCoinFacade(ctx, CoinHandler)

	c := routes.NewCoinRoute(CoinFacade)
	c.CoinRoutes(server)
	server.Run(":" + "7000")
}
