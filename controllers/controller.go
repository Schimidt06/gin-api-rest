package controllers

import (
	"github.com/Schimidt06/gin-api-rest/database"
	"github.com/Schimidt06/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

// ExibeTodosAlunos responde a requisição com a lista de todos os alunos
func ExibeTodosAlunos(c *gin.Context) {
	// Cria uma slice de Aluno para armazenar os dados vindos do banco
	var alunos []models.Aluno

	// Busca todos os registros no banco de dados usando a instância DB
	database.DB.Find(&alunos)

	// Retorna a lista de alunos em formato JSON com status 200 (OK)
	c.JSON(200, alunos)
}

// Saudacao exibe uma mensagem de boas-vindas personalizada
func Saudacao(c *gin.Context) {
	// Captura o parâmetro 'nome' vindo da URL (ex: /:nome)
	nome := c.Params.ByName("nome")

	// Retorna um JSON com a mensagem de saudação
	c.JSON(200, gin.H{
		"message": "Olá, seja bem-vindo(a) " + nome,
	})
}

// CriaNovoAluno recebe os dados em JSON e cria um novo aluno no banco de dados
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	// Valida se o JSON recebido bate com a estrutura do modelo
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Cria o aluno no banco de dados
	database.DB.Create(&aluno)

	// Retorna o aluno criado
	c.JSON(201, aluno)
}
