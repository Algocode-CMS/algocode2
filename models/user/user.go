package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Login        string `gorm:"unique"`
	PasswordHash string
	Email        string `gorm:"unique"`

	Verified bool
	IsAdmin  bool

	PasswordResetCode *string

	UserData `gorm:"embedded"`
}

type UserData struct {
	Surname     *string
	Name        *string
	FathersName *string
	BirthDate   *time.Time

	Country    *string
	Region     *string
	City       *string
	School     *string
	SchoolCity *string
	Grade      *string

	Telegram    *string
	PhoneNumber *string

	CFLogin *string

	PCMSLogin    *string
	PCMSPassword *string

	EjudgeLogin    *string
	EjudgePassword *string
	EjudgeID       int
}
