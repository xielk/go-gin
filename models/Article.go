// article.go

package models

import (
	"time"
)

// GaArticle represents the ga_articles table in the database
type Article struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;not null"`
	Title        string    `gorm:"size:128;default:null;comment:'标题'"`
	Mark         string    `gorm:"size:128;default:null;comment:'标签'"`
	Desc         string    `gorm:"size:256;default:null;comment:'描述'"`
	Cover        string    `gorm:"size:256;default:null;comment:'图片地址'"`
	Video        string    `gorm:"size:128;default:null;comment:'视频地址'"`
	Source       string    `gorm:"size:64;default:null;comment:'来源'"`
	Content      string    `gorm:"type:text;not null;comment:'内容'"`
	CreatedAt    time.Time `gorm:"column:create_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:'建立时间'" json:",omitempty"`
	UpdatedAt    time.Time `gorm:"column:update_at;type:timestamp;default:null;comment:'更新时间'"  json:",omitempty"`
	ReadCount    int       `gorm:"type:int;default:0;comment:'阅读数'"`
	CommentCount int       `gorm:"type:int;default:0;comment:'评论数'"`
	StarCount    int       `gorm:"type:int;default:0;comment:'点赞数'"`
	Sort         int16     `gorm:"type:smallint;default:255;comment:'排序'"`
	Status       bool      `gorm:"type:tinyint;default:1;comment:'是否审核'"`
	Type         int       `gorm:"type:tinyint;default:1;comment:'文章类型，文章或视频'"`
	Comments     []Comment `gorm:"foreignKey:ArticleID"`
	Stars        []Star    `gorm:"foreignKey:ArticleID"`
}

func (Article) TableName() string {
	return "ga_articles"
}

// UpdateArticleInput defines the input structure for updating an article.
type UpdateArticleInput struct {
	Title        string    `json:"title"`
	Mark         string    `json:"mark"`
	Description  string    `json:"description"`
	Video        string    `json:"video_url"`
	Source       string    `json:"source"`
	Content      string    `json:"content"`
	UpdatedAt    time.Time `json:"updated_at"`
	ReadCount    int       `json:"read_count"`
	CommentCount int       `json:"comment_count"`
	StarCount    int       `json:"star_count"`
	Sort         int16     `json:"sort"`
	Status       bool      `json:"status"`
}

type CreateArticleInput struct {
	Title        string    `json:"title"`
	Mark         string    `json:"mark"`
	Description  string    `json:"description"`
	Video        string    `json:"video_url"`
	Source       string    `json:"source"`
	Content      string    `json:"content"`
	UpdatedAt    time.Time `json:"updated_at"`
	ReadCount    int       `json:"read_count"`
	CommentCount int       `json:"comment_count"`
	StarCount    int       `json:"star_count"`
	Sort         int16     `json:"sort"`
	Status       bool      `json:"status"`
}
