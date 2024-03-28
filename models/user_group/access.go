package user_group

import (
	"github.com/Algocode-CMS/algocode2/models/common"
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/choices"
)

type AccessEmbed struct {
	AccessForUnregistered choices.Select[common.AccessRight]
	AccessForRegistered   choices.Select[common.AccessRightNil]
	AccessForGroup        choices.Select[common.AccessRightNil]
	GroupID               *uint // TODO: Google how to add foreign key here

	// TODO: Add admin user access
}
