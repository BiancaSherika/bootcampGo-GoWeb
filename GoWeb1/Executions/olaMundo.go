package executations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OlaMundo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Olá Bianca Claro",
	})

}
