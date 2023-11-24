package mainController

import (
	"apiGolang/models/category"
	"apiGolang/models/detailIncomeAndExpense"
	"apiGolang/models/incomeAndExpense"
	"apiGolang/models/repositories"
)

func init() {

	repositories.IncomeAndExpensesRepo = incomeAndExpense.GetRepo()
	repositories.CategoryRepo = category.GetRepo()
	repositories.DetailIncomeAndExpenseRepo = detailIncomeAndExpense.GetRepo()
}
