package controllers

import (
	"go-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "github.com/swaggo/files"
)

type HomeController struct {
	DB *gorm.DB
	BaseController
}

func (ac *HomeController) Index(c *gin.Context) {

	var (
		menus      []models.Menu
		page       int
		pageSize   int
		totalItems int64
	)
	// 获取分页参数并处理错误
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err = strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	// 计算总菜单数
	if err := ac.DB.Model(&models.Menu{}).Count(&totalItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法计算总菜单数"})
		return
	}

	// 分页查询菜单
	if err := ac.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&menus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法查询菜单"})
		return
	}

	// 构造响应结构
	response := ac.NewListResponse(page, pageSize, totalItems, menus)
	c.JSON(http.StatusOK, response)
}
