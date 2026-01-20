package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/Schimidt06/gin-api-rest/database"
	"github.com/Schimidt06/gin-api-rest/models"
	"github.com/gin-gonic/gin"
)

// ExibeTodosAlunos responde a requisição com a lista de todos os alunos
func ExibeTodosAlunos(c *gin.Context) {
	// Cria uma slice de Aluno para armazenar os dados vindos do banco
	var alunos []models.Aluno

	// Inicializa a query base
	query := database.DB

	// Filtro por Nome (parcial e case-insensitive)
	if nome := c.Query("nome"); nome != "" {
		query = query.Where("nome ILIKE ?", "%"+nome+"%")
	}

	// Filtro por CPF (exato)
	if cpf := c.Query("cpf"); cpf != "" {
		query = query.Where("cpf = ?", cpf)
	}

	// Filtro por RG (exato)
	if rg := c.Query("rg"); rg != "" {
		query = query.Where("rg = ?", rg)
	}

	// Busca os registros no banco de dados com os filtros aplicados
	// Paginação
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var total int64
	// Conta o total de registros ANTES de aplicar o limite, mas COM os filtros
	query.Model(&models.Aluno{}).Count(&total)

	// Aplica offset e limit e busca
	query.Offset(offset).Limit(limit).Find(&alunos)

	// Calcula total de páginas
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	// Retorna JSON estruturado com metadados
	c.JSON(200, gin.H{
		"data":       alunos,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": totalPages,
	})
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

// BuscaAlunoPorID localiza um aluno pelo ID
func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")   // Captura o ID enviado na URL
	database.DB.First(&aluno, id) // Busca no banco o aluno com este ID

	// Verifica se o aluno foi encontrado (ID diferente de 0)
	if aluno.ID == 0 {
		c.JSON(404, gin.H{
			"Not Found": "Aluno não encontrado",
		})
		return
	}

	// Retorna os dados do aluno encontrado
	c.JSON(200, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"message": "Aluno deletado com sucesso"})
}

// EditaAluno localiza um aluno pelo ID e atualiza seus dados
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	// Tenta encontrar o aluno no banco
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(404, gin.H{"error": "Aluno não encontrado"})
		return
	}

	// Faz o bind do JSON recebido para a struct do aluno (atualiza os campos)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Salva as alterações no banco de dados usando UpdateColumns para evitar problemas com zero values se necessário,
	// ou Save se quisermos sobrescrever tudo. UpdateColumns é mais seguro para PATCH parcial.
	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(200, aluno)
}
