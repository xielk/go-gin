// article.go

package models

import (
	"time"
	
)

type Article struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"column:title"`
	Description   string    `gorm:"column:desc"`
	Content       string    `gorm:"column:content"`
	CoverImageURL string    `gorm:"column:cover_image_url"`
	DeletedOn     uint      `gorm:"column:deleted_on"`
	Status        uint8     `gorm:"column:status"`
	CreatedAt     time.Time `gorm:"column:create_at"`
	UpdatedAt     time.Time `gorm:"column:update_at"`

	Comments []Comment `gorm:"foreignKey:ArticleID"`
	Stars    []Star    `gorm:"foreignKey:ArticleID"`
}

func (Article) TableName() string {
	return "tb_articles"
}

// CreateArticleInput defines the input structure for creating an article.
type CreateArticleInput struct {
    Title         string `json:"title"`
    Description   string `json:"description"`
    Content       string `json:"content"`
    CoverImageURL string `json:"cover_image_url"`
}

// UpdateArticleInput defines the input structure for updating an article.
type UpdateArticleInput struct {
    Title         string `json:"title"`
    Description   string `json:"description"`
    Content       string `json:"content"`
    CoverImageURL string `json:"cover_image_url"`
}

