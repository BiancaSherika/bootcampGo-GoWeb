package executations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OlaMundo(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Olá Bianca Claro",
		})
	})
}
