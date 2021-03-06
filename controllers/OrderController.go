package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
// OrderController ini merupakan kontrak yang dipakai ketika dipanggil di app.go
type OrderController interface {
	CreateOrder(ctx *gin.Context)
	GetAllData(ctx *gin.Context)
	DispatchData(ctx *gin.Context)
	DestroyData(ctx *gin.Context)
}

// di orderController ini service yang akan dibutuhkan
type orderController struct {
	// block code dimasukan  kesini yang akan menghandle service
}

// NewOrderControlle ini merupakan function yang dibuat untuk menyatukan interface dan struct
func NewOrderControlle() OrderController {
	return &orderController{}
}

func (c *orderController) Create(ctx *gin.Context) {
	// ini kaya response di express
	ctx.JSON(http.StatusOK, gin.H {
		"message": "ini routingan create order",
	})
}

func (c *orderController) GetAllData(ctx *gin.Context) {
	// ini kaya response di express
	ctx.JSON(http.StatusOK, gin.H {
		"message": "ini routingan get all order",
	})
}

func (c *orderController) DispatchData(ctx *gin.Context) {
	// ini kaya response di express
	ctx.JSON(http.StatusOK, gin.H {
		"message": "ini routingan patch edit order",
	})
}

func (c *orderController) DestroyData(ctx *gin.Context) {
	// ini kaya response di express
	ctx.JSON(http.StatusOK, gin.H {
		"message": "ini routingan delete order",
	})
}