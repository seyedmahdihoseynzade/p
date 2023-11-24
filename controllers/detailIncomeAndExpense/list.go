package detailIncomeAndExpense

import (
	"apiGolang/apiSchema/detailIncomeAndExpense"
	"apiGolang/controllers/mainController"
	"apiGolang/models/repositories"
	"context"
	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	req := detailIncomeAndExpense.List{}
	mainController.ParseBody(ctx, &req)
	res, code, err := repositories.DetailIncomeAndExpenseRepo.List(context.Background(), req)
	if err != nil {
		return mainController.Error(ctx, req, code, err)
	}
	return mainController.Response(ctx, res)
}
