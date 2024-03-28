package course

import (
	"fmt"
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/file"
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/priority"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model

	Name        string
	Description string

	Photo file.Holder

	TelegramID string
	Courses    []*TeacherInCourse
}

func (t *Teacher) ToString() string {
	return fmt.Sprintf("[%d] %s", t.ID, t.Name)
}

type TeacherInCourse struct {
	ID        uint `gorm:"primarykey"`
	CourseID  uint
	TeacherID uint

	Note string
	priority.Embed
}
