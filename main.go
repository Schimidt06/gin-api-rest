package main

import (
	"github.com/gin-gonic/gin" // Importa o framework Gin, que ajuda a criar aplicações web (APIs) de forma rápida
)

// ExibeTodosAlunos é a "função controladora" (handler) que vai responder quando alguém acessar a rota de alunos
func ExibeTodosAlunos(c *gin.Context) {
	// c.JSON envia uma resposta no formato JSON para quem fez a requisição
	// 200 é o código HTTP de "Sucesso" (OK)
	// gin.H é um atalho para criar um mapa de dados (chave: valor)
	c.JSON(200, gin.H{
		"id":    1,      // Retorna o ID do aluno
		"nome":  "João", // Retorna o nome do aluno
		"idade": 25,     // Retorna a idade do aluno
	})
}

func main() {
	// Cria uma nova instância do servidor Gin com configurações padrão (logger e recovery)
	r := gin.Default()

	// Define uma rota do tipo GET.
	// Quando alguém acessar "http://localhost:8080/alunos", a função ExibeTodosAlunos será executada
	r.GET("/alunos", ExibeTodosAlunos)

	// Inicia o servidor na porta padrão (8080)
	// O servidor ficará rodando e "escutando" requisições até você pará-lo
	r.Run()
}
