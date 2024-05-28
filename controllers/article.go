package controllers

import (
	"errors"
	"fmt"
	"go-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "github.com/swaggo/files"
)

type ArticleController struct {
	DB *gorm.DB
	BaseController
}

// GetArticles godoc
// @Summary 获取文章列表
// @Description 获取文章列表，支持分页查询
// @Tags Articles
// @Accept  json
// @Produce  json
// @Param page query int false "页码 (默认为1)"
// @Param page_size query int false "每页显示数量 (默认为10)"
// @Success 200 {object} []models.Article
// @Failure 500 {object} ErrorResponse
// @Router /articles [get]
func (ac *ArticleController) GetArticles(c *gin.Context) {
	var (
		articles   []models.Article
		page       int
		pageSize   int
		totalItems int64
	)
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("page_size", "10"))
	// 使用预加载获取文章及其相关的评论和星级
	query := ac.DB

	//query.Preload("Stars.User") //列出点赞的用户属性
	// 计算总文章数
	query.Model(&models.Article{}).Count(&totalItems)
	// 分页查询文章
	query.Order("sort ASC").Offset((page - 1) * pageSize).Limit(pageSize).Find(&articles)
	// 构造响应结构
	response := ac.NewListResponse(page, pageSize, totalItems, articles)
	c.JSON(http.StatusOK, OK("Success", response))
}

// GetArticle godoc
// @Summary 获取文章详情 type=1表示文章，type=2表示视频。视频在列表点击直接播放
// @Description 根据文章ID获取文章详情
// @Tags Articles
// @Accept  json
// @Produce  json
// @Param id query int true "文章ID"
// @Success 200 {object} models.Article
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /article [get]
func (ac *ArticleController) GetArticle(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("Fetching article with ID:", id)

	// 查询文章
	var article models.Article
	if err := ac.DB.Preload("Stars").Preload("Comments").First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, ErrorResponse{Message: "Article not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, OK("Success", article))
}

func (ac *ArticleController) CreateArticle(c *gin.Context) {
	var input models.CreateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input data"})
		return
	}

	// 创建文章
	article := models.Article{
		Title:   input.Title,
		Content: input.Content,
		// 设置其他字段
	}

	if err := ac.DB.Create(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create article"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (ac *ArticleController) UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	// 查询文章
	var article models.Article
	if err := ac.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Article not found"})
		return
	}

	var input models.UpdateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid input data"})
		return
	}

	// 更新文章
	article.Title = input.Title
	article.Content = input.Content
	// 更新其他字段

	if err := ac.DB.Save(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to update article"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	// 查询文章
	var article models.Article
	if err := ac.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Article not found"})
		return
	}

	// 删除文章
	if err := ac.DB.Delete(&article).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to delete article"})
		return
	}

	c.JSON(http.StatusOK, "删除成功")
}
