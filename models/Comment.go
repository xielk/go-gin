// comment.go

package models

import "time"

type Comment struct {
	ID        uint         `gorm:"primaryKey"`
	ArticleID uint         `gorm:"column:aid"`
	UserID    uint         `gorm:"column:uid"`
	Title     string       `gorm:"column:title"`
	Content   string       `gorm:"column:content"`
	CreatedAt time.Time    `gorm:"column:create_at"`
	UpdatedAt time.Time    `gorm:"column:update_at"`
	User      ResponseUser `gorm:"foreignKey:UserID"` // 使用ResponseUser类型
}

func (Comment) TableName() string {
	return "ga_article_comments"
}
