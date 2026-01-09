package controllers

import (
	"github.com/Schimidt06/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"message": "Olá, seja bem-vindo(a) " + nome,
	})
}
