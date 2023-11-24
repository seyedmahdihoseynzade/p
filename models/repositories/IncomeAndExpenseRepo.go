package repositories

import (
	"apiGolang/apiSchema/IncomeAndExpense"
	"apiGolang/models/incomeAndExpense/dataModel"
	"context"
)

var IncomeAndExpensesRepo IncomeAndExpenses

type IncomeAndExpenses interface {
	IncomeAndExpensesList(ctx context.Context) (res IncomeAndExpense.IncomeAndExpensesList, code int, err error)
	Create(ctx context.Context, req IncomeAndExpense.Create) (code int, err error)
	GetByID(ID, userID int64) dataModel.IncomeAndExpense
}
