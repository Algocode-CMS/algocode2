package course

import (
	"fmt"
	"github.com/Algocode-CMS/algocode2/models/common"
	"github.com/Algocode-CMS/algocode2/models/user_group"
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/priority"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Label string `gorm:"unique"`

	// Will we keep previous design?
	Title       string
	Subtitle    string
	CoursesText string `gorm:"default:'Курсы'"`

	MainLinks  []*GroupMainLink
	RightLinks []*GroupRightLink
	Groups     []*user_group.UserGroup

	common.JudgeData
	user_group.AccessEmbed
}

type GroupMainLink struct {
	ID      uint `gorm:"primarykey"`
	GroupID uint

	Text string
	Link string

	priority.Embed
}

type GroupRightLink struct {
	ID      uint `gorm:"primarykey"`
	GroupID uint
	common.Link
}

func (g *Group) AfterFind(*gorm.DB) error {
	priority.Sort(g.MainLinks)
	priority.Sort(g.RightLinks)
	return nil
}

func (g *Group) ToString() string {
	return fmt.Sprintf("[%d] %s (%s)", g.ID, g.Label, g.Title)
}
