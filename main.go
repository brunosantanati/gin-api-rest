package main

import (
	"github.com/brunosantanati/api-go-gin/database"
	"github.com/brunosantanati/api-go-gin/models"
	"github.com/brunosantanati/api-go-gin/routes"
)

/*
Rodar: go run main.go
Acessar: http://localhost:8080/
*/

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Bruno Sant' Ana", CPF: "00000000000", RG: "320000000"},
		{Nome: "Ana", CPF: "11111111111", RG: "330000000"},
	}

	routes.HandleRequests()
}
