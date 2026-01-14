package models

import "gorm.io/gorm"

// Aluno representa a estrutura de dados de um estudante no sistema
type Aluno struct {
	gorm.Model
	Nome  string `json:"nome"`  // Nome do aluno
	CPF   string `json:"cpf"`   // CPF do aluno
	RG    string `json:"rg"`    // RG do aluno
	Idade int    `json:"idade"` // Idade do aluno
	Sexo  string `json:"sexo"`  // Sexo do aluno
}
