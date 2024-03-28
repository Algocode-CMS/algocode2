package common

import (
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/file"
	"github.com/Algocode-CMS/algocode2/pkg/admin/models/priority"
)

type Link struct {
	Text string
	Link string

	Hidden bool
	NewTab bool
	File   file.Holder

	priority.Embed
}
