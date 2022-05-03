package main

import (
	"fmt"
	"go-rest-api-mux/database"
	"go-rest-api-mux/models"
	"go-rest-api-mux/routes"
)

func main() {
	models.Personalidades = []models.Personalidade {
		{Id: 1, Nome: "Silvio Santos", Historia: "Apresentador de TV"},
		{Id: 2, Nome: "Ronaldo", Historia: "Jogador de Futebol"},
		{Id: 3, Nome: "Ayrton Senna", Historia: "Piloto de F1"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleRequest()
}

