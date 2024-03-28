package config

type LinksBlock struct {
	// Title stores title for current block. Can be nil
	Title string

	Links []*BlockLink
}

type BlockLink struct {
	Title string

	// Link is either to some other admin page, or to parts in path.
	// Path and PageIDs can use following types:
	// - string
	// - PathID
	// - ModelField
	Page    *PageConfig
	PageIDs []interface{}

	Path []interface{}
}
