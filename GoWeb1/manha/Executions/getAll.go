package executations

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetAll(r *gin.Engine) {
	r.GET("/usuarios", func(c *gin.Context) {
		users, err := os.ReadFile("usuarios.json")
		if err != nil {
			log.Fatal("Erro ao abrir arquivo json", err)
			c.JSON(http.StatusNotFound, gin.H{"Erro": "Erro ao abrir o arquivo JSON"})
		}

		var pessoas []Usuario
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
