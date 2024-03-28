package config

type PageConfig struct {
	// Parent page. Can be nil only for main page.
	Parant *PageConfig

	// Label to add to url. Should be unique for every son of a parent
	Label string

	// MainModel stores pointer to a struct for which the page is built. Can be nil
	//
	// If MainModel is nil, the url of current page will be: <parent_url>/<label>/
	// Else the url of current page will be: <parent_url>/<label>/<structure_id>/
	MainModel interface{}

	// Title for current page
	Title string

	// HideNavPath can hide parent pages in nav path
	HideNavPath bool

	// NavPathTitle is used for this page in nav path
	NavPathTitle string

	// Blocks stores pointers to blocks that should be inside current page.
	//
	// Content can be one of the following:
	// - LinksBlock
	// - ModelListBlock
	// - ModelFieldsBlock
	// - TextBlock
	// - ChildInlineBlock
	Blocks []interface{}
}
