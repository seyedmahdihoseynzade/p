package route

import (
	"apiGolang/controllers/category"
	"github.com/gofiber/fiber/v2"
)

var categoryR = map[string]string{
	"create": "/category/create",
	"list":   "/category/list",
}

func SetupCategoryRoutesRoute(app *fiber.App) map[string]string {
	app.Post(categoryR["list"], category.List)
	app.Post(categoryR["create"], category.Create)

	return categoryR
}
