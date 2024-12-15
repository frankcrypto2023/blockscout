package qngplex

import (
	"qngplex/pkg/dbmodels"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(conf *Config) {
	var err error
	db, err = gorm.Open(postgres.Open(conf.PostgresUrl), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	dbmodels.InitBlockTable(db)
	dbmodels.InitTxTable(db)
}

func GetDB() *gorm.DB {
	return db
}

func StopDB() {
	if db != nil {
	}
}
