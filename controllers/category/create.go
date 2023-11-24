package category

import (
	"apiGolang/apiSchema/category"
	"apiGolang/controllers/mainController"
	"apiGolang/models/repositories"
	"context"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	req := category.Create{}
	mainController.ParseBody(ctx, &req)
	code, err := repositories.CategoryRepo.Create(context.Background(), req)
	if err != nil {
		return mainController.Error(ctx, req, code, err)
	}
	return mainController.Response(ctx, nil)
}
