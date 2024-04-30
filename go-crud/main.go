package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/article", controllers.ArticlesCreate)
	r.PUT("/article/:id", controllers.ArticlesUpdate)
	r.DELETE("/article/:id", controllers.ArticlesDelete)

	r.GET("/articles/:limit/:offset", controllers.ArticlesShow)
	r.GET("/article/:id", controllers.ArticlesId)
	r.Run() // listen and serve on 0.0.0.0:8080
}
