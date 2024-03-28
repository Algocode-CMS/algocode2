package user_group

import (
	"fmt"
	"gorm.io/gorm"
)

type UserGroup struct {
	gorm.Model
	CourseID uint

	Name      string
	ShortName string

	Users []*UserInGroup
}

func (ug *UserGroup) ToString() string {
	return fmt.Sprintf("[%d] %s (%s)", ug.ID, ug.Name, ug.ShortName)
}

// TODO: Add UserGroup actions (e.g. register for contest)
