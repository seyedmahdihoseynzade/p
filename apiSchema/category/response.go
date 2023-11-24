package category

type CategoryList struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
