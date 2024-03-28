package course

import (
	"database/sql"
	"fmt"
	"github.com/Algocode-CMS/algocode2/models/common"
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/choices"
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/file"
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/priority"
	"gorm.io/gorm"
	"time"
)

type Contest struct {
	gorm.Model

	CourseID uint
	Course   Course `gorm:"foreignKey:CourseID"`

	Date           time.Time
	Title          string
	Statements     file.Holder
	ShowStatements bool

	Type  choices.Select[common.ContestType]
	Judge choices.Select[common.ContestJudge]

	ContestID       string
	DurationSeconds string

	priority.Embed

	Additional ContestAdditional `gorm:"embedded"`
	Links      []*ContestLink
}

type ContestAdditional struct {
	Display         bool
	ExternalGroupID string
	OtherLink       string
	PcmsStandings   string
	ExternalXMLLink string
	ScoreLatest     bool
	ScoreFinished   bool
	EnableStartTime bool
	StartTime       sql.NullTime
	ContestInfo     map[string]string `gorm:"serializer:json"`
}

type ContestLink struct {
	ID        uint `gorm:"primarykey"`
	ContestID uint
	common.Link
}

func (c *Contest) AfterFind(*gorm.DB) error {
	priority.Sort(c.Links)
	return nil
}

func (c *Contest) ToString() string {
	return fmt.Sprintf("[%d] %s", c.ID, c.Title)
}
