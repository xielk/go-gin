// controllers/user.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	 // jwt "github.com/dgrijalva/jwt-go"
	// "go-gin/models"
	// "github.com/swaggo/files"
)

type UserController struct{
	BaseController
}

func (uc *UserController) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user",
	})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create user",
	})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update user",
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete user",
	})
}
