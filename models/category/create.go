package category

import (
	"apiGolang/apiSchema/category"
	"apiGolang/models/category/dataModel"
	"context"
)

func (repo *Repository) Create(ctx context.Context, req category.Create) (int, error) {
	err := repo.db().Create(ctx, dataModel.Category{
		UserID:   1,
		Name:     req.Name,
		ParentID: req.ParentCategory,
	})
	return 200, err
}
