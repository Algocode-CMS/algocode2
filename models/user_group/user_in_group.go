package user_group

import "github.com/Algocode-CMS/algocode2/models/user"

type UserInGroup struct {
	ID uint `gorm:"primaryKey"`

	UserGroupID uint
	UserID      uint
	User        user.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`

	Name      *string
	CFLogin   *string
	EjudgeID  *string
	PCMSLogin *string
}
