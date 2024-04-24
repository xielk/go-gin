// comment.go

package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	ArticleID uint      `gorm:"column:aid"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:create_at"`
	UpdatedAt time.Time `gorm:"column:update_at"`
}

func (Comment) TableName() string {
	return "tb_article_comments"
}
