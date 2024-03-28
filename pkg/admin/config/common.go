package config

// PathID is used for getting ids in current path. Ids are numbered from parent to child
type PathID int

// ModelField can be used to obtain field of loaded models
type ModelField struct {
	// Idx denotes index of model in page
	// 0 - index of model for current page (if page has model)
	// >0 - index of model in inlined models lists (1 - first inline, 2 - second level inline and so on)
	// <0 - index of model of parent pages (-1 - parent, -2 - parent of parent and so on)
	Idx int

	// Field denotes name of field of current model
	Field string
}
