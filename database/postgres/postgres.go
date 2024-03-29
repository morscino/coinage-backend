package postgres

import (
	"github.com/morscino/gigo/utility/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect(database config.DatabaseConfig) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=" + database.User + " password=" + database.Password + " dbname=" + database.Name + " sslmode=" + database.SSLMode,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		//log.Error("Could not connect to database : %v", err)
		panic(err.Error())
	}

	return db
}
