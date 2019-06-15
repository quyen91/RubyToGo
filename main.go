package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"authen-api/controllers"
)


func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/users", controllers.CreateUser)
		v1.GET("/users", controllers.ListUsers)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}