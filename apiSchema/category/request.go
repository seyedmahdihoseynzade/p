package category

type Create struct {
	Name           string `json:"name"`
	ParentCategory int64  `json:"parentCategory"`
}

type CategoryListReq struct {
	CategoryID int64 `json:"categoryID"`
}
