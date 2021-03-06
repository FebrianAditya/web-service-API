package main

import (
	"github.com/febrianaditya/web-service-API/config"
	controller "github.com/febrianaditya/web-service-API/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDBConnection() // ini untuk memanggil config yang sudah dibuat di folder config
	
	userController controller.UserController = controller.NewUserController()
)

func main() {
	r := gin.Default()
	
	userRoutes := r.Group("user") // --> ini sama kaya di folder Routes file index.js untk end pointnya
	{ // ---> ini harus di bawah curly bracketnya kaya di VueRoutes dan AngularRoutes
		// hit endpointnya mirip express, kasih controller
		userRoutes.POST("/login", userController.Login) // ---> ini kaya route di folder Routes di file yang selain index.js, methode dan end point nya dibuat di sini
		userRoutes.POST("/register", userController.Register)
	}

	orderRoutes := r.Group("order") // --> ini sama kaya di folder Routes file index.js untk end pointnya
	{ // ---> ini harus di bawah curly bracketnya kaya di VueRoutes dan AngularRoutes
		// hit endpointnya mirip express, kasih controller
		orderRoutes.POST("/", orderController.CreateOrder) // ---> ini kaya route di folder Routes di file yang selain index.js, methode dan end point nya dibuat di sini
		orderRoutes.GET("/", orderController.GetAllData)
		orderRoutes.DISPATCH("/", orderController.DispatchData)
		orderRoutes.DELETE("/", orderController.DestroyData)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	
}