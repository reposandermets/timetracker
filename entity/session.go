package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// https://gist.github.com/WangYihang/7d43d70db432ff8f3a0a88425bfca7f2
type User struct {
	ID       uuid.UUID `json:"id" gorm:"column:id;primary_key;type:char(36);"`
	Name     string    `json:"name" binding:"required" gorm:"type:varchar(255)"`
	Sessions []Session `gorm:"ForeignKey:UserID"`
}

type Session struct {
	ID        uuid.UUID `json:"id" gorm:"column:id;primary_key;type:char(36);"`
	UserID    uuid.UUID `json:"user_id" binding:"required"`
	Name      string    `json:"name"  gorm:"type:varchar(255)"`
	Status    string    `json:"status"  gorm:"type:varchar(255)"`
	StartedAt time.Time `json:"started_at"`
	StopperAt time.Time `json:"stopper_at"`
	EndedAt   time.Time `json:"ended_at"`
	Seconds   float64   `json:"seconds"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
