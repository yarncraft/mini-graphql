package main

import (
	"github.com/yarncraft/mini-graphql/src/models"
	connection "github.com/yarncraft/mini-graphql/src/db"
)

func main() {
	db := connection.GetConnection()
	defer connection.Close()

	db.AutoMigrate(&models.Post{}, &models.User{})
}