package main

import (
	. "apiGolang/services/core/route"
	log "code.ts.co.ir/gaplib/golib/logger"
	. "code.ts.co.ir/gaplib/golib/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var routes map[string]string

func setupRoute(app *fiber.App) {
	app.Get("/buildTime/core", func(ctx *fiber.Ctx) error {
		log.GetLogger().Info("build at "+buildTime, zap.String("ip", ctx.Get("x-real-ip")))
		return ctx.SendString("build at " + buildTime)
	})
	routes = MergeMaps(
		SetupCapitalManagementRoutesRoute(app),
		SetupCategoryRoutesRoute(app),
		SetupDetailIncomeAndExpensesRoute(app),
	)
}
