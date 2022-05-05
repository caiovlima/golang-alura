package main

import (
	"go-rest-api-gin/database"
	"go-rest-api-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
