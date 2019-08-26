package db

import (
	"sync"
    "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var instance *gorm.DB
var once sync.Once

func InitializeDB(env string) *gorm.DB {
	db, err := gorm.Open("postgres", "postgres://postgres:docker@localhost:5432?sslmode=disable")
  	if err != nil {
    	panic("failed to connect database")
	  }

	return db
}

func GetConnection() *gorm.DB {
	env := "dev"
	once.Do(func() {
		instance = InitializeDB(env)
		if env != "prod" {
			instance.LogMode(true)
		}
	})
	return instance
}

func Close() {
	if instance == nil {
		return
	}
	instance.Close()
}