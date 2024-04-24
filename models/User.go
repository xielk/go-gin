// user.go

package models

import "time"

type User struct {
	ID            uint       `gorm:"primaryKey"`
	Username      string     `gorm:"column:user_name"`
	Password      string     `gorm:"column:pass"`
	Realname      string     `gorm:"column:real_name"`
	Phone         string     `gorm:"column:phone"`
	Status        int        `gorm:"column:status"`
	Remark        string     `gorm:"column:remark"`
	LastLoginTime *time.Time `gorm:"column:last_login_time"`
	LastLoginIP   string     `gorm:"column:last_login_ip"`
	LoginTimes    int        `gorm:"column:login_times"`
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "tb_users"
}
