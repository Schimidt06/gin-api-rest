package database

import (
	"log"

	"github.com/Schimidt06/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB // Instância da conexão com o banco de dados
	err error    // Variável para tratamento de erros
)

// ConectaComBancoDeDados inicializa a conexão com o PostgreSQL usando GORM
func ConectaComBancoDeDados() {
	// String de conexão com as credenciais e configurações do banco
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	// Abre a conexão com o banco de dados
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		// Encerra a aplicação caso não consiga conectar
		log.Panic("Erro ao conectar com banco de dados")
	}

	// Realiza a migração automática, criando a tabela 'alunos' baseada na struct Aluno
	DB.AutoMigrate(&models.Aluno{})
}
