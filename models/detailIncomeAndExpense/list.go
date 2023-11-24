package detailIncomeAndExpense

import (
	"apiGolang/apiSchema/detailIncomeAndExpense"
	"apiGolang/models/repositories"
	"code.ts.co.ir/gaplib/golib/cast"
	"code.ts.co.ir/gaplib/golib/timeUtils"
	"context"
	"errors"
	ptime "github.com/yaa110/go-persian-calendar"
	"time"
)

func (repo *Repository) List(ctx context.Context, req detailIncomeAndExpense.List) (res detailIncomeAndExpense.IncomeAndExpensesList, code int, err error) {
	parentID := req.ParentID
	if parentID == 0 {
		return res, 200, errors.New("آیدی موضو اصلی اجباری است")
	}
	i := repositories.IncomeAndExpensesRepo.GetByID(parentID, 1)
	if i.ID == 0 {
		return res, 200, errors.New("جزيیات درخواسنی نامعتبر است")
	}

	if i.Amount > 0 {
		return res, 200, errors.New("جزيیات فقط برای برداشت ها موجود است")
	}

	res.Month = ptime.New(time.Unix(i.Date, 0)).Month().String() + cast.ForceToString(ptime.New(time.Unix(i.Date, 0)).Year())
	res.Name = i.Title
	res.Budget = -i.Amount
	list := repo.db().GetList(parentID)
	for j := range list {
		t := "برداشت"
		c := repositories.CategoryRepo.GatByID(ctx, 1, list[j].CategoryID)
		res.IncomeAndExpensesList = append(res.IncomeAndExpensesList, detailIncomeAndExpense.IncomeAndExpenses{
			Day:         ptime.New(time.Unix(timeUtils.ConvertNanoSecToSec(list[j].ID), 0)).Weekday().String(),
			Time:        cast.ForceToString(ptime.New(time.Unix(timeUtils.ConvertNanoSecToSec(list[j].ID), 0)).Hour()) + ":" + cast.ForceToString(ptime.New(time.Unix(timeUtils.ConvertNanoSecToSec(list[j].ID), 0)).Minute()),
			Title:       list[j].Title,
			Description: list[j].Description,
			Amount:      list[j].Amount,
			Type:        t,
			Category:    c.Name,
		})
	}
	return res, 200, nil
}
