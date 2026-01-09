package main

import (
	"github.com/Schimidt06/gin-api-rest/models"
	"github.com/Schimidt06/gin-api-rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{
			Nome:  "João",
			Idade: 25,
			Sexo:  "Masculino",
			CPF:   "12345678901",
			RG:    "123456789",
		},
		{
			Nome:  "Maria",
			Idade: 20,
			Sexo:  "Feminino",
			CPF:   "10987654321",
			RG:    "109876543",
		},
	}
	routes.HandleRequests()
}
