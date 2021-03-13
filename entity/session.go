package entity

import "time"

type User struct {
	ID   uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Name string `json:"name" binding:"required" gorm:"type:varchar(256)"`
}

type Session struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Status    string    `json:"status" binding:"required" gorm:"type:varchar(10)"`
	User      User      `json:"user" binding:"required" gorm:"foreignkey:UserID"`
	UserID    uint64    `json:"-"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
