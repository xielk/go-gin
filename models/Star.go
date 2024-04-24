// star.go

package models

import "time"

type Star struct {
	ID        uint      `gorm:"primaryKey"`
	ArticleID uint      `gorm:"column:aid"`
	UserID    uint      `gorm:"column:uid"`
	CreatedAt time.Time `gorm:"column:create_at"`
	UpdatedAt time.Time `gorm:"column:update_at"`
	Username  string    `gorm:"-"`        // 不映射到数据库，用于存储用户名
	User User `gorm:"foreignKey:UserID"` // 定义与 User 模型的关联
}

func (Star) TableName() string {
	return "tb_article_stars"
}


