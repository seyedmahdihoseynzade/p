package dataSources

import (
	"apiGolang/models/incomeAndExpense/dataModel"
	"apiGolang/models/incomeAndExpense/dataSources/vitessDS"
)

type IncomeAndExpenseDBDS interface {
	GetList(userID int64) []dataModel.IncomeAndExpense
	Create(m dataModel.IncomeAndExpense) error
	GetByID(ID, userID int64) dataModel.IncomeAndExpense
	TotalBalance(userID int64) float32
}

func GetIncomeAndExpenseDBDS() IncomeAndExpenseDBDS {
	return vitessDS.GetVitessDS()
}
