package category

import (
	"apiGolang/models/category/dataModel"
	"apiGolang/models/category/dataSources"
	"context"
	"sync"
)

type Repository struct {
	dbDS dataSources.CategoryDBDS
}

var (
	repo   *Repository
	doOnce sync.Once
)

func initRepo() {
	repo = &Repository{}
}

func GetRepo() *Repository {
	doOnce.Do(initRepo)
	return repo
}

func (repo *Repository) db() dataSources.CategoryDBDS {
	if repo.dbDS == nil {
		repo.dbDS = dataSources.GetCategoryDBDS()
	}

	return repo.dbDS
}

func (repo *Repository) GatByID(ctx context.Context, userID, ID int64) dataModel.Category {
	return repo.db().GatByID(ctx, userID, ID)
}
