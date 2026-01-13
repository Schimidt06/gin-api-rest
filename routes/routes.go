package routes

import (
	"github.com/Schimidt06/gin-api-rest/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// HandleRequests define todas as rotas da API
func HandleRequests() {
	// Cria uma instância padrão do Gin com logger e recovery middleware
	r := gin.Default()

	// Configura o CORS para permitir requisições de qualquer origem
	r.Use(cors.Default())

	// Define a rota GET para /alunos que lista todos os alunos
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	// Define a rota GET com parâmetro dinâmico :nome para saudação
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)

	// Inicia o servidor na porta padrão (8080)
	r.Run()
}
