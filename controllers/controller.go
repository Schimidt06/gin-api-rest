package controllers

import (
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":    1,
		"nome":  "João",
		"idade": 25,
	})
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"message": "Olá, seja bem-vindo(a) " + nome,
	})
}
