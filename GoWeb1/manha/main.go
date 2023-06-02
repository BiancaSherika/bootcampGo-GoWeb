package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	Id          int
	Nome        string
	Sobrenome   string
	Email       string
	Idade       int
	Altura      float64
	Ativo       bool
	DataCriacao string
}

func olaMundo(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Olá Bianca Claro",
		})
	})
}

func getAll(r *gin.Engine) {
	r.GET("/usuarios", func(c *gin.Context) {
		users, err := os.ReadFile("usuarios.json")
		if err != nil {
			log.Fatal("Erro ao abrir arquivo json", err)
			c.JSON(http.StatusNotFound, gin.H{"Erro": "Erro ao abrir o arquivo JSON"})
		}

		var pessoas []usuario
		resultado := json.Unmarshal([]byte(users), &pessoas)

		if resultado != nil {
			log.Fatal("Erro ao abrir arquivo json", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Erro": "Erro ao ler o arquivo JSON"})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": pessoas,
		})
	})
}

func main() {
	r := gin.Default()
	olaMundo(r)
	getAll(r)
	r.Run()
}
