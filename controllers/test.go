package controllers

import (
	"net/http"

	"go-gin/config"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	cfg config.Config
	BaseController
}

func (tc *TestController) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Test",
	})
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /test/helloworld [get]
func (tc *TestController) HelloWorld(c *gin.Context)  {
	c.JSON(http.StatusOK,"helloworld")
}

// func (tc *TestController) Config(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"server":   tc.cfg.Server,
// 		"database": tc.cfg.Database,
// 	})
// }
