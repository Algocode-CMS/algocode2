package config

type ModelListBlock struct {
	// Model stores pointer to a struct for which the list is built
	Model interface{}

	// QueryConds stores conditions for filtering in gorm format.
	// Following types can be used:
	// - string
	// - PathID
	// - ModelField
	QueryConds []interface{}

	// DisplayedFields stores field names to display (by default only ID will be displayed)
	DisplayedFields []string

	// MaxOnPage stores maximum number of models on page,
	MaxOnPage int

	// Page denotes page for which the link should be created to edit model
	Page *PageConfig

	// PageIDs are ids that will be inserted in page path. Can be of following types:
	// - string
	// - PathID
	// - ModelField
	// - ModelID (id of current model in the list)
	PageIDs []interface{}

	// TODO: Add query for obtaining model
}

type ModelID uint
