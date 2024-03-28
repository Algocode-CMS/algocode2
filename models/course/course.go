package course

import (
	"fmt"
	"github.com/Algocode-CMS/algocode2/models/common"
	"github.com/Algocode-CMS/algocode2/models/user_group"
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/priority"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model

	Label string `gorm:"index:,unique,composite:group_course"`

	GroupID uint  `gorm:"index:,unique,composite:group_course"`
	Group   Group `gorm:"foreignKey:GroupID"`

	Title        string
	Subtitle     string
	ContestsText string `gorm:"default:'Курсы'"`

	common.JudgeData
	user_group.AccessEmbed

	Contests   []*Contest
	RightLinks []*CourseRightLink
	Teachers   []*TeacherInCourse
}

type CourseRightLink struct {
	ID       uint `gorm:"primarykey"`
	CourseID uint
	common.Link
}

func (c *Course) AfterFind(d *gorm.DB) error {
	priority.Sort(c.Contests)
	priority.Sort(c.RightLinks)
	priority.Sort(c.Teachers)
	return nil
}

func (c *Course) ToString() string {
	return fmt.Sprintf("[%d] %s (%s)", c.ID, c.Label, c.Title)
}
