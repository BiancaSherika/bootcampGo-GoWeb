package main

import (
	"github.com/BiancaSherika/bootcampGo-GoWeb/GoWeb/cmd/server/handler"
	"github.com/BiancaSherika/bootcampGo-GoWeb/GoWeb/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := handler.NewUser(service)

	r := gin.Default()
	us := r.Group("/users")
	us.GET("/", u.GetAll())
	us.POST("/", u.PostUser())
	us.PUT("/:id", u.PutUser())
	us.PATCH("/:id", u.PatchName())
	us.DELETE("/:id", u.DeleteUser())
	r.Run()
}
