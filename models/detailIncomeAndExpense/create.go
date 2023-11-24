package detailIncomeAndExpense

import (
	"apiGolang/apiSchema/detailIncomeAndExpense"
	"apiGolang/models/detailIncomeAndExpense/dataModel"
	"apiGolang/models/repositories"
	"context"
	"errors"
)

func (repo *Repository) Create(ctx context.Context, req detailIncomeAndExpense.Create) (code int, err error) {

	c := repositories.CategoryRepo.GatByID(ctx, 1, req.Category)
	if c.ID == 0 {
		return 200, errors.New("دسته بندی وجود ندارد")
	}
	if c.ParentID == 0 {
		return 200, errors.New("دسته بندی نامعتبر است (برای هزینه های جزئی فقط دسته بندی های زیر مجموعه باید استفاده شود)")
	}

	if req.ParentIncomeAndExpenseID == 0 {
		return 200, errors.New("آیدی موضوع اصلی اجباری است")
	}
	i := repositories.IncomeAndExpensesRepo.GetByID(req.ParentIncomeAndExpenseID, 1)
	if i.ID == 0 {
		return 200, errors.New("موضوع اصلی نامعتبر است")
	}

	if i.CategoryID != c.ParentID {
		return 200, errors.New("دسته بندی نامعتبر است")
	}

	if i.Amount >= 0 {
		return 200, errors.New("جزيیات فقط برای برداشت ها موجود است")
	}

	if req.Amount >= 0 {
		return 200, errors.New("ققط برداشت را ثبت کنید")
	}

	sumOfAmount := repo.db().SumOfAmount(i.ID)
	if -i.Amount+req.Amount+sumOfAmount < 0 {
		return 200, errors.New("موجودی کافی نیست")
	}

	err = repo.db().Create(dataModel.DetailIncomeAndExpense{
		Title:       req.Title,
		Description: req.Description,
		Amount:      req.Amount,
		CategoryID:  req.Category,
		PrentID:     req.ParentIncomeAndExpenseID,
	})
	return 0, err
}
