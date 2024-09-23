package database

import (
	"fmt"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/container"
	"github.com/ali-mahdavi-bn/service-site/src/organization/adapter/data_model"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var clients = map[string]func(dsn string) gorm.Dialector{
	"sqlite":   sqlite.Open,
	"postgres": postgres.Open,
}

type Config struct {
	Debug       bool
	AutoMigrate bool
}

func InitDB(config *Config) {
	dns := ":memory:"
	dns = "gorm.db"
	driver := "sqlite"

	if envDns, ok := os.LookupEnv("DATABASE_HOST"); ok {
		dns = envDns
	}
	if envDriver, ok := os.LookupEnv("DATABASE_DRIVER"); ok {
		driver = envDriver
	}
	var db *gorm.DB
	if client, ok := clients[driver]; ok {
		var err error
		db, err = gorm.Open(client(dns), &gorm.Config{SkipDefaultTransaction: true})
		if err != nil {
			panic("failed to connect database")
		}
	}
	if config.Debug {
		db = db.Debug()
	}
	fmt.Println(&data_model.OrderLines{})
	if config.AutoMigrate {
		err := db.AutoMigrate(
			&data_model.OrderLines{},
			&data_model.Batches{},
		)
		if err != nil {
			panic("failed to automigrate")
		}
	}
	container.DB = db
}
