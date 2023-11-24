package category

import (
	"apiGolang/apiSchema/category"
	"apiGolang/models/category/dataModel"
	"context"
)

func (repo *Repository) List(ctx context.Context, req category.CategoryListReq) (category.CategoryList, int, error) {
	l := []dataModel.Category{}
	if req.CategoryID == 0 {
		l = repo.db().GetMainCategory(ctx, 1)
	} else {
		l = repo.db().GetSubCategory(ctx, 1, req.CategoryID)
	}

	c := category.CategoryList{}
	for i := range l {
		c.Categories = append(c.Categories, category.Category{
			ID:   l[i].ID,
			Name: l[i].Name,
		})
	}
	return c, 200, nil
}
