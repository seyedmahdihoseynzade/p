package dataSources

import (
	"apiGolang/models/detailIncomeAndExpense/dataModel"
	"apiGolang/models/detailIncomeAndExpense/dataSources/vitessDS"
)

type DetailIncomeAndExpenseDBDS interface {
	GetList(prentID int64) []dataModel.DetailIncomeAndExpense
	Create(d dataModel.DetailIncomeAndExpense) error
	SumOfAmount(parentID int64) float32
}

func GetDetailIncomeAndExpenseDBDS() DetailIncomeAndExpenseDBDS {
	return vitessDS.GetVitessDS()
}
