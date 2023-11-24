package incomeAndExpense

import (
	"apiGolang/apiSchema/IncomeAndExpense"
	"apiGolang/models/incomeAndExpense/dataModel"
	"apiGolang/models/repositories"
	"code.ts.co.ir/gaplib/golib/timeUtils"
	"context"
	"errors"
)

func (repo *Repository) Create(ctx context.Context, req IncomeAndExpense.Create) (code int, err error) {

	c := repositories.CategoryRepo.GatByID(ctx, 1, req.Category)
	if c.ID == 0 {
		return 200, errors.New("دسته بندی وجود ندارد")
	}
	if c.ParentID != 0 {
		return 200, errors.New("دسته بندی نامعتبر است (برای هزینه های کلی فقط دسته بندی های اصلی باید استفاده شود)")
	}

	err = repo.db().Create(dataModel.IncomeAndExpense{
		UserID:      1,
		Date:        timeUtils.GetThisMonthStartTime().Unix(),
		Title:       req.Title,
		Description: req.Description,
		Amount:      req.Amount,
		CategoryID:  req.Category,
	})
	return 0, err
}
