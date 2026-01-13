package main

import (
	"github.com/Schimidt06/gin-api-rest/database"
	"github.com/Schimidt06/gin-api-rest/models"
	"github.com/Schimidt06/gin-api-rest/routes"
)

// main é o ponto de entrada da aplicação
func main() {
	// Conecta ao banco de dados PostgreSQL
	database.ConectaComBancoDeDados()

	// Cria dados iniciais (Seed) apenas para exemplo
	var aluno models.Aluno
	if err := database.DB.First(&aluno).Error; err != nil {
		alunos := []models.Aluno{
			{Nome: "João", CPF: "12345678901", RG: "123456789", Idade: 25, Sexo: "Masculino"},
			{Nome: "Maria", CPF: "10987654321", RG: "109876543", Idade: 20, Sexo: "Feminino"},
		}
		database.DB.Create(&alunos)
	}

	// Inicializa e houve as requisições HTTP
	routes.HandleRequests()
}
