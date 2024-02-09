package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func InitFiber() *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      viper.GetString("APP_NAME"),
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
