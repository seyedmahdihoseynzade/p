package mainController

import (
	"code.ts.co.ir/gaplib/golib/logger"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

type Request interface{}

func ParseBody(ctx *fiber.Ctx, request interface{}) (int, error) {
	var (
		err        error
		statusCode int
	)

	logger := log.GetLogger()
	body := ctx.Body()

	if len(body) != 0 {
		err = jsoniter.Unmarshal(body, request)
		if err != nil {
			logger.Error("body unmarshal error", zap.Error(err))
			return 400, err
		}
	}

	return statusCode, err
}

func prepareBody(ctx *fiber.Ctx, body interface{}) error {
	logger := log.GetLogger()

	var (
		data []byte
		err  error
	)

	data, err = jsoniter.Marshal(body)
	if err != nil {
		logger.Error("body marshal error", zap.Error(err))
		return err
	}
	ctx.Set("Content-type", "application/json")
	_, _ = ctx.Write(data)

	return nil
}

func Response(ctx *fiber.Ctx, data interface{}) error {
	ctx.Status(200)
	ctx.Set("Content-type", "application/json")
	return prepareBody(ctx, data)
}

func Error(ctx *fiber.Ctx, req interface{}, code int, err error) error {
	ctx.Status(code)

	x := struct {
		Req interface{} `json:"request"`
		Err interface{} `json:"errorText"`
	}{}

	x.Req = req
	x.Err = err.Error()
	return prepareBody(ctx, x)
}
