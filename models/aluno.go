package models

import "gorm.io/gorm"

// Aluno representa a estrutura de dados de um estudante no sistema
type Aluno struct {
	gorm.Model
	Nome  string `json:"nome" binding:"required"`       // Nome é obrigatório
	CPF   string `json:"cpf" binding:"required,len=11"` // CPF obrigatório e deve ter 11 chars
	RG    string `json:"rg" binding:"required,min=7"`   // RG obrigatório, mín 7 chars
	Idade int    `json:"idade" binding:"required"`      // Idade obrigatória
	Sexo  string `json:"sexo" binding:"required"`       // Sexo obrigatório
}
