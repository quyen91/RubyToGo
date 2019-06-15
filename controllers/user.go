package controllers

import (
	"authen-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

var CreateUser = func(c *gin.Context) {
	user := &models.User{
		Name: c.Query("name"),
		Email: c.Query("email"),
	}

	resp := user.Create()

	c.JSON(200, gin.H(resp))
}

var ListUsers = func(c *gin.Context) {


	fmt.Print(models.GetListUsers())
	c.JSON(200, gin.H{
		"users": models.GetListUsers(),
	})
}

var ShowUser = func(c *gin.Context) {

}
