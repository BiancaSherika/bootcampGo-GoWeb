package main

import (
	executations "github.com/BiancaSherika/bootcampGo-GoWeb/GoWeb1/manha/Executions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	executations.OlaMundo(r)
	executations.GetAll(r)
	r.Run()
}
