package repositories

import (
	"apiGolang/apiSchema/detailIncomeAndExpense"
	"context"
)

var DetailIncomeAndExpenseRepo DetailIncomeAndExpense

type DetailIncomeAndExpense interface {
	List(ctx context.Context, req detailIncomeAndExpense.List) (res detailIncomeAndExpense.IncomeAndExpensesList, code int, err error)
	Create(ctx context.Context, req detailIncomeAndExpense.Create) (code int, err error)
}
