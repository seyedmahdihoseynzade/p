package incomeAndExpense

import (
	"apiGolang/apiSchema/IncomeAndExpense"
	"apiGolang/models/incomeAndExpense/dataModel"
	"apiGolang/models/repositories"
	"code.ts.co.ir/gaplib/golib/cast"
	"code.ts.co.ir/gaplib/golib/timeUtils"
	"code.ts.co.ir/gaplib/golib/utils"
	"context"
	ptime "github.com/yaa110/go-persian-calendar"
	"sort"
	"time"
)

func (repo *Repository) IncomeAndExpensesList(ctx context.Context) (res IncomeAndExpense.IncomeAndExpensesList, code int, err error) {

	list := repo.db().GetList(1)
	m := make(map[int64][]dataModel.IncomeAndExpense)
	dates := make([]int64, 0)
	for i := range list {
		dates = append(dates, list[i].Date)
		m[list[i].Date] = append(m[list[i].Date], list[i])
	}

	dates = utils.UniqueArray(dates)
	sort.Slice(dates, func(i, j int) bool {
		return dates[i] > dates[j]
	})

	res.TotalBalance = repo.TotalBalance(1)
	for i := range dates {
		sort.Slice(m[dates[i]], func(k, j int) bool {
			return m[dates[i]][k].ID > m[dates[i]][j].ID
		})
		income := IncomeAndExpense.IncomeAndExpensesPerMonth{
			Month: ptime.New(time.Unix(dates[i], 0)).Month().String() + cast.ForceToString(ptime.New(time.Unix(dates[i], 0)).Year()),
		}
		for j := range m[dates[i]] {
			t := ""
			if m[dates[i]][j].Amount < 0 {
				t = "برداشت"
			} else if m[dates[i]][j].Amount > 0 {
				t = "واریز"
			}
			c := repositories.CategoryRepo.GatByID(ctx, 1, m[dates[i]][j].CategoryID)
			income.IncomeAndExpensesList = append(income.IncomeAndExpensesList, IncomeAndExpense.IncomeAndExpenses{
				ID:          m[dates[i]][j].ID,
				Day:         ptime.New(time.Unix(timeUtils.ConvertNanoSecToSec(m[dates[i]][j].ID), 0)).Weekday().String(),
				Time:        cast.ForceToString(ptime.New(time.Unix(timeUtils.ConvertNanoSecToSec(m[dates[i]][j].ID), 0)).Hour()) + ":" + cast.ForceToString(ptime.New(time.Unix(timeUtils.ConvertNanoSecToSec(m[dates[i]][j].ID), 0)).Minute()),
				Title:       m[dates[i]][j].Title,
				Description: m[dates[i]][j].Description,
				Amount:      m[dates[i]][j].Amount,
				Type:        t,
				Category: IncomeAndExpense.Category{
					ID:   c.ID,
					Name: c.Name,
				},
			})
		}
		res.IncomeAndExpensesList = append(res.IncomeAndExpensesList, income)
	}

	return res, 200, nil
}
