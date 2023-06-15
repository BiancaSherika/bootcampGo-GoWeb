package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/BiancaSherika/bootcampGo-GoWeb/GoWeb/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string
	LastName     string
	Email        string
	Age          int
	Height       float64
	Active       bool
	CreationDate string
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func CheckError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		panic(err)
	}
}

func (*User) CheckToken(ctx *gin.Context) {
	token := ctx.Request.Header.Get("TOKEN")
	if token != os.Getenv("TOKEN") {
		err := errors.New("token inválido")
		CheckError(ctx, err)
	}
}

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.CheckToken(ctx)

		u, err := c.service.GetAll()
		CheckError(ctx, err)
		ctx.JSON(200, u)
	}
}

func (c *User) PostUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.CheckToken(ctx)

		var req request
		err := ctx.Bind(&req)
		CheckError(ctx, err)

		u, err := c.service.PostUser(req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CreationDate)
		CheckError(ctx, err)
		ctx.JSON(200, u)
	}
}

func (c *User) PutUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.CheckToken(ctx)

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		CheckError(ctx, err)

		var req request
		err = ctx.ShouldBindJSON(&req)
		CheckError(ctx, err)

		if req.Name == "" {
			err := errors.New("nome é obrigatório")
			CheckError(ctx, err)
		}

		if req.Email == "" {
			err := errors.New("email é obrigatório")
			CheckError(ctx, err)
		}

		if req.LastName == "" {
			err := errors.New("sobrenome é obrigatório")
			CheckError(ctx, err)
		}

		if req.Age == 0 {
			err := errors.New("idade é obrigatório")
			CheckError(ctx, err)
		}

		if req.Height == 0 {
			err := errors.New("altura é obrigatório")
			CheckError(ctx, err)
		}

		if req.CreationDate == "" {
			err := errors.New("data de criação é obrigatório")
			CheckError(ctx, err)
		}

		u, err := c.service.PutUser(int(id), req.Name, req.LastName, req.Email, req.Age, req.Height, req.Active, req.CreationDate)
		CheckError(ctx, err)

		ctx.JSON(200, u)
	}

}

func (c *User) PatchName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.CheckToken(ctx)

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		CheckError(ctx, err)

		var req request
		err = ctx.ShouldBindJSON(&req)
		CheckError(ctx, err)

		if req.Name == "" {
			err := errors.New("nome é obrigatório")
			CheckError(ctx, err)
		}

		u, err := c.service.PatchName(int(id), req.Name)
		CheckError(ctx, err)

		ctx.JSON(200, u)
	}
}

func (c *User) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.CheckToken(ctx)

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		CheckError(ctx, err)

		err = c.service.DeleteUser(int(id))
		CheckError(ctx, err)

		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O usuário %d foi removido", id)})
	}
}
