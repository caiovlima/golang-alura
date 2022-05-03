package main

import (
	"fmt"
	"go-rest-api-mux/models"
	"go-rest-api-mux/routes"
)

func main() {
	models.Personalidades = []models.Personalidade {
		{Nome: "Silvio Santos", Historia: "Apresentador de TV"},
		{Nome: "Ronaldo", Historia: "Jogador de Futebol"},
		{Nome: "Ayrton Senna", Historia: "Piloto de F1"},
	}

	fmt.Println("Iniciando servidor com Go")
	routes.HandleRequest()
}

