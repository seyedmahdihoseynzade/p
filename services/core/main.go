package main

import (
	service "code.ts.co.ir/gaplib/golib/gapServices"
)

var buildTime = "unknown"

func main() {
	app := service.New("core")
	setupRoute(app)
	if err := app.Listen(":7575"); err != nil {
		panic(err)
	}
}
