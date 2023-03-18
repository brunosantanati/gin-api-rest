package main

import "github.com/gin-gonic/gin"

/*
Rodar: go run main.go
Acessar: http://localhost:8080/
*/

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Bruno Sant' Ana",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos)
	r.Run()
}
