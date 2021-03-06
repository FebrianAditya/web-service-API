package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController ini merupakan kontrak yang akan dipakai ketika dipanggil di app.go
type UserController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

// di userController ini service yang akan dibutuhkan
type userController struct {
	// block code dimasukan  kesini yang akan menghandle service
}

// NewUserController ini merupakan function yang dibuat untuk menyatukan interface dan struct
func NewUserController() UserController {
	return &userController{}
}

func (c *userController) Login(ctx *gin.Context) {
	// ini kaya response di express
	ctx.JSON(http.StatusOK, gin.H {
		"message": "ini routingan login",
	})
}

func (c *userController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"message": "ini routingan Register",
	})
}