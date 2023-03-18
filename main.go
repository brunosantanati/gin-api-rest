package main

import (
	"github.com/brunosantanati/api-go-gin/database"
	"github.com/brunosantanati/api-go-gin/routes"
)

/*
Rodar: go run main.go
Acessar: http://localhost:8080/
*/

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
