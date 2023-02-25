package main

import (
	"example/demo1/controllers"
	"example/demo1/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariable()
	initializer.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
