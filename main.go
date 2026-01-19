package main

import (
	"github.com/Schimidt06/gin-api-rest/database"
	"github.com/Schimidt06/gin-api-rest/routes"
)

// main é o ponto de entrada da aplicação
func main() {
	// Conecta ao banco de dados PostgreSQL
	database.ConectaComBancoDeDados()

	// Inicializa e ouve as requisições HTTP
	routes.HandleRequests()
}
