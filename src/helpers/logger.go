package helpers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetLoggerField(ctx *fiber.Ctx, category string, data interface{}) logrus.Fields {
	return logrus.Fields{
		"service":         "ping",
		"http_method":     ctx.Method(),
		"ip_address":      ctx.IP(),
		"hostname":        ctx.Hostname(),
		"method":          fmt.Sprintf(`[%s] %s`, ctx.Method(), ctx.OriginalURL()),
		"uri":             ctx.OriginalURL(),
		"body":            ctx.Body(),
		"category":        category,
		"additional_info": data,
	}
}
