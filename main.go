package main

import (
	"github.com/Schimidt06/gin-api-rest/database"
	"github.com/Schimidt06/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
