package main

import (
	executations "github.com/BiancaSherika/bootcampGo-GoWeb/GoWeb1/Executions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", executations.OlaMundo)
	r.GET("/usuarios/GetAll", executations.GetAll)
	r.GET("/usuarios/:id", executations.GetById)
	r.Run()
}
