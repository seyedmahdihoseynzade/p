package route

import (
	"apiGolang/controllers/IncomeAndExpense"
	"github.com/gofiber/fiber/v2"
)

var capitalManagementRoutes = map[string]string{
	"incomeAndExpensesList": "/IncomeAndExpense/incomeAndExpenses/list",
	"create":                "/IncomeAndExpense/incomeAndExpenses/create",
}

func SetupCapitalManagementRoutesRoute(app *fiber.App) map[string]string {
	app.Post(capitalManagementRoutes["incomeAndExpensesList"], IncomeAndExpense.List)
	app.Post(capitalManagementRoutes["create"], IncomeAndExpense.Create)

	return capitalManagementRoutes
}
