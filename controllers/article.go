package controllers

import (
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

// GetArticles 获取文章列表
// @Summary 获取所有文章列表，包括评论和星级信息
// @Description 获取所有文章列表，包括评论和星级信息
// @Tags Articles
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {array} models.Article
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
	query := ac.DB.Preload("Comments").Preload("Stars").Preload("Stars.User")
	// 计算总文章数
	query.Model(&models.Article{}).Count(&totalItems)
	// 分页查询文章
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&articles)
	// 构造响应结构
	response := ac.NewListResponse(page, pageSize, totalItems, articles)
	c.JSON(http.StatusOK, response)
}

// GetArticle 获取单篇文章详情
// @Summary 获取单篇文章详情
// @Description 获取单篇文章详情
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} models.Article
// @Failure 400 {object} ErrorResponse
// @Router /article/{id} [get]
func (ac *ArticleController) GetArticle(c *gin.Context) {
	id := c.Param("id")

	// 查询文章
	var article models.Article
	if err := ac.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Article not found"})
		return
	}

	c.JSON(http.StatusOK, article)
}

// CreateArticle 创建文章
// @Summary 创建文章
// @Description 创建文章
// @Tags Articles
// @Accept json
// @Produce json
// @Param input body models.CreateArticleInput true "文章信息"
// @Success 200 {object} models.Article
// @Failure 400 {object} ErrorResponse
// @Router /article [post]
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

// UpdateArticle 更新文章
// @Summary 更新文章
// @Description 更新文章
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Param input body models.UpdateArticleInput true "文章信息"
// @Success 200 {object} models.Article
// @Failure 400 {object} ErrorResponse
// @Router /article/{id} [put]
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

// DeleteArticle 删除文章
// @Summary 删除文章
// @Description 删除文章
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {string} string "删除成功"
// @Failure 400 {object} ErrorResponse
// @Router /article/{id} [delete]
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
