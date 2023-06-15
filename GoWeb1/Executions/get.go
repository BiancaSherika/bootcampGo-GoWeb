package executations

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	users, err := ioutil.ReadFile("usuarios.json")
	if err != nil {
		tratamentoErro(err, c)
	}

	var pessoas []Usuario
	err = json.Unmarshal([]byte(users), &pessoas)
	if err != nil {
		tratamentoErro(err, c)
	}

	c.JSON(http.StatusOK, pessoas)
}

func GetById(c *gin.Context) {

	users, err := ioutil.ReadFile("usuarios.json")
	if err != nil {
		tratamentoErro(err, c)
	}

	var pessoas []Usuario
	err = json.Unmarshal([]byte(users), &pessoas)
	if err != nil {
		tratamentoErro(err, c)
	}

	filtroId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		tratamentoErro(err, c)
	}

	for _, pessoa := range pessoas {
		if pessoa.Id == filtroId {
			c.JSON(http.StatusOK, gin.H{
				"Usuário:": pessoa,
			})
		}
	}
}

func tratamentoErro(err error, c *gin.Context) {
	log.Fatal("Erro ao abrir arquivo json", err)
	c.JSON(http.StatusBadRequest, gin.H{
		"Erro": err.Error(),
	})
}
