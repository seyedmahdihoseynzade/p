package incomeAndExpense

import (
	"apiGolang/models/incomeAndExpense/dataModel"
	"apiGolang/models/incomeAndExpense/dataSources"
	"sync"
)

type Repository struct {
	dbDS dataSources.IncomeAndExpenseDBDS
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

func (repo *Repository) db() dataSources.IncomeAndExpenseDBDS {
	if repo.dbDS == nil {
		repo.dbDS = dataSources.GetIncomeAndExpenseDBDS()
	}

	return repo.dbDS
}

func (repo *Repository) GetByID(ID, userID int64) dataModel.IncomeAndExpense {
	return repo.db().GetByID(ID, userID)
}

func (repo *Repository) TotalBalance(userID int64) float32 {
	return repo.db().TotalBalance(userID)
}
