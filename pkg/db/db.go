package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"web-app/config"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *config.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn))
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
