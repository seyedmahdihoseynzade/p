package category

import (
	"apiGolang/apiSchema/category"
	"apiGolang/controllers/mainController"
	"apiGolang/models/repositories"
	"context"
	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	req := category.CategoryListReq{}
	mainController.ParseBody(ctx, &req)
	res, code, err := repositories.CategoryRepo.List(context.Background(), req)
	if err != nil {
		return mainController.Error(ctx, req, code, err)
	}
	return mainController.Response(ctx, res)
}
