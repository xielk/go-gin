// article.go

package models

import (
	"time"
)

type Menu struct {
	ID       uint       `gorm:"primaryKey;autoIncrement"`
	Title    string     `gorm:"type:varchar(64)"`
	Icon     string     `gorm:"type:varchar(128)"`
	URL      string     `gorm:"type:varchar(256)"`
	AppID    string     `gorm:"type:varchar(64)"`
	Type     int        `gorm:"type:int;default:1;comment:'1 微信小程序\n2 支付宝小程序\n3 H5\n4 APP'"`
	CreateAt time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdateAt *time.Time `gorm:"type:timestamp"`
	MsgCount int        `gorm:"type:int;default:0;comment:'消息数'"`
	CatName  string     `gorm:"type:varchar(128)"`
	Sort     int        `gorm:"type:int;default:255;comment:'排序'"`
}

func (Menu) TableName() string {
	return "ga_menus"
}

type CreateMenuInput struct {
	Title    string `json:"title" binding:"required,max=64"`
	Icon     string `json:"icon" binding:"max=128"`
	URL      string `json:"url" binding:"max=256"`
	AppID    string `json:"appid" binding:"max=64"`
	Type     int    `json:"type" binding:"required,oneof=1 2 3 4"`
	MsgCount int    `json:"msg_count" binding:"max=4"`
	Sort     int    `json:"sort" binding:"max=255"`
	CatName  string `json:"cat_name" binding:"max=32"`
}

// UpdateMenuInput 用于更新菜单的输入数据
type UpdateMenuInput struct {
	Title    *string    `json:"title" binding:"omitempty,max=64"`
	Icon     *string    `json:"icon" binding:"omitempty,max=128"`
	URL      *string    `json:"url" binding:"omitempty,max=256"`
	AppID    *string    `json:"appid" binding:"omitempty,max=64"`
	Type     *int       `json:"type" binding:"omitempty,oneof=1 2 3 4"`
	MsgCount *int       `json:"msg_count" binding:"omitempty,max=4"`
	UpdateAt *time.Time `json:"update_at"`
	Sort     int        `json:"sort" binding:"max=255"`
	CatName  string     `json:"cat_name" binding:"max=32"`
}
