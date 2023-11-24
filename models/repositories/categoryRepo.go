package repositories

import (
	"apiGolang/apiSchema/category"
	"apiGolang/models/category/dataModel"
	"context"
)

var CategoryRepo Category

type Category interface {
	Create(ctx context.Context, req category.Create) (int, error)
	List(ctx context.Context, req category.CategoryListReq) (category.CategoryList, int, error)
	GatByID(ctx context.Context, userID, ID int64) dataModel.Category
}
