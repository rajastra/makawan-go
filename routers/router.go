package routers

import (
	"io"
	"os"
	"tubes/controllers"
	"tubes/midlewares"

	"github.com/gin-gonic/gin"
)

func setLogger() {
	file, _ := os.Create("dts.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
func GetRouter() *gin.Engine {
	setLogger()
	router := gin.New()
	//engine.Use(Logger(), Recovery())
	router.Use(gin.Recovery(), midlewares.Logger(), midlewares.GenerateIdUnix())

	//http://localhost:8080/api/v1
	api := router.Group("/api/v1")
	{
		//GET http://localhost:8080/api/v1/health
		api.GET("/health", controllers.HealthCheck)

		//  Product
		// POST http://localhost:8080/api/v1/product
		api.POST("/product", controllers.CreateProduct)
		// GET http://localhost:8080/api/v1/product
		api.GET("/product", controllers.GetAllProduct)

		//http://localhost:8080/api/v1/users
		groupuser := api.Group("/users")

		//POST http://localhost:8080/api/v1/users/registrasi
		groupuser.POST("/registrasi", controllers.Registrasi)
		//POST http://localhost:8080/api/v1/users/login
		groupuser.POST("/login", controllers.Login)

		//http://localhost:8080/api/v1/transacation    // harus menggunakan token jwt ada use / midleware
		grouptrans := api.Group("/transacation").Use(midlewares.JWTAuth())
		{
			//POST http://localhost:8080/api/v1/transacation/orders  & body json
			grouptrans.POST("/orders", controllers.CreateOrder)

			//GET http://localhost:8080/api/v1/transacation/orders/12345
			grouptrans.GET("/orders/:orderID", controllers.GetOrderBy)

			//PUT http://localhost:8080/api/v1/transacation/orders/12345  & body json
			grouptrans.PUT("/orders/:orderID", controllers.UpdateOrder)

			//DELETE http://localhost:8080/api/v1/transacation/orders/12345
			grouptrans.DELETE("/orders/:orderID", controllers.DeleteOrder)
		}

	}

	return router
}
