package IncomeAndExpense

import (
	"apiGolang/controllers/mainController"
	"apiGolang/models/repositories"
	"context"
	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	res, code, err := repositories.IncomeAndExpensesRepo.IncomeAndExpensesList(context.Background())
	if err != nil {
		return mainController.Error(ctx, nil, code, err)
	}
	return mainController.Response(ctx, res)
}
