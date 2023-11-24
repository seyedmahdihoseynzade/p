package route

import (
	"apiGolang/controllers/detailIncomeAndExpense"
	"github.com/gofiber/fiber/v2"
)

var detailIncomeAndExpensesRoutes = map[string]string{
	"list":   "/detailIncomeAndExpenses/list",
	"create": "/detailIncomeAndExpenses/create",
}

func SetupDetailIncomeAndExpensesRoute(app *fiber.App) map[string]string {
	app.Post(detailIncomeAndExpensesRoutes["list"], detailIncomeAndExpense.List)
	app.Post(detailIncomeAndExpensesRoutes["create"], detailIncomeAndExpense.Create)

	return detailIncomeAndExpensesRoutes
}
