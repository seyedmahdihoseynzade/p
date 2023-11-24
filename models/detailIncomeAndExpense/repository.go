package detailIncomeAndExpense

import (
"sync"
"apiGolang/models/detailIncomeAndExpense/dataSources"
)

type Repository struct {
	dbDS            dataSources.DetailIncomeAndExpenseDBDS
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

func (repo *Repository) db() dataSources.DetailIncomeAndExpenseDBDS{
	if repo.dbDS == nil {
		repo.dbDS = dataSources.GetDetailIncomeAndExpenseDBDS()
	}

	return repo.dbDS
}
