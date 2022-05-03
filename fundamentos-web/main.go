package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
} 

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto {
		{"Camiseta", "Azul, bem bonita", 39, 5},
		{"Tenis", "Preto, bem bonito", 50, 2},
		{"Fone", "Branco de boa qualidade", 10, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}