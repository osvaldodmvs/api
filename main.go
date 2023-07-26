package main

import (
	"github.com/gin-gonic/gin"
	"github.com/osvaldodmvs/api/controllers"
	"github.com/osvaldodmvs/api/initializers"
	"github.com/osvaldodmvs/api/middleware"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.GET("/validate", middleware.AuthMiddleware, controllers.Validate) //just for testing

	//CRUD
	r.POST("/products", middleware.AuthMiddleware, controllers.CreateProduct)
	r.GET("/products", middleware.AuthMiddleware, controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductById)
	r.PUT("/products/:id", middleware.AuthMiddleware, controllers.UpdateProductById)
	r.DELETE("/products/:id", middleware.AuthMiddleware, controllers.DeleteProductById)

	r.Run()
}
