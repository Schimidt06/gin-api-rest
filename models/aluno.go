package models

type Aluno struct {
	Nome  string `json:"nome"`
	CPF   string `json:"cpf"`
	RG    string `json:"rg"`
	Idade int    `json:"idade"`
	Sexo  string `json:"sexo"`
}

var Alunos []Aluno
