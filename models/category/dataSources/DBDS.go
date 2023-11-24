package dataSources

import (
	"apiGolang/models/category/dataModel"
	"apiGolang/models/category/dataSources/vitessDS"
	"context"
)

type CategoryDBDS interface {
	Create(ctx context.Context, category dataModel.Category) error
	GatByID(ctx context.Context, userID, ID int64) dataModel.Category
	GetMainCategory(ctx context.Context, userID int64) []dataModel.Category
	GetSubCategory(ctx context.Context, userID int64, prentID int64) []dataModel.Category
}

func GetCategoryDBDS() CategoryDBDS {
	return vitessDS.GetVitessDS()
}
