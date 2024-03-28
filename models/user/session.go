package user

import (
	"time"
)

type Session struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	User       User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	ValidUntil time.Time
}
