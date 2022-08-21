package main

import (
	"github.com/rodrigoadao/simple-go-gin-api/database"
	"github.com/rodrigoadao/simple-go-gin-api/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}
