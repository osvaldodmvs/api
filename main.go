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

	r.LoadHTMLGlob("templates/**/*")
	r.Static("/assets", "./assets")

	r.GET("/", controllers.Home)
	r.GET("/signup", controllers.SignUpPage)
	r.POST("/signup", controllers.SignUp)
	r.GET("/login", controllers.LoginPage)
	r.POST("/login", controllers.Login)

	//r.GET("/validate", middleware.AuthMiddleware, controllers.Validate) //just for testing

	//CRUD
	r.GET("/product", middleware.AuthMiddleware, controllers.CreateProductPage)
	r.POST("/product", middleware.AuthMiddleware, controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/product/:id", controllers.GetProductById)
	r.PUT("/product/edit/:id", middleware.AuthMiddleware, controllers.UpdateProductById)
	r.DELETE("/product/delete/:id", middleware.AuthMiddleware, controllers.DeleteProductById)

	/* Works for html but not ideal for security
	r.GET("/product/edit/:id", middleware.AuthMiddleware, controllers.UpdateProductById)
	r.GET("/product/delete/:id", middleware.AuthMiddleware, controllers.DeleteProductById)
	*/

	r.Run()
}
