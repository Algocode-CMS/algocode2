package config

type ModelFieldsBlock struct {
	Title string

	HiddenFields []string
	// If ShownFields not nil, only following fields will be shown
	ShownFields []string
	// ID is always read only and should not be included here
	ReadOnlyFields []string
	HideAllFields  bool

	ExcludedEmbeddings    []string
	HiddenEmbeddings      []string
	HideAllEmbeddings     bool
	SpecializedEmbeddings []*ModelFieldsBlock
}
