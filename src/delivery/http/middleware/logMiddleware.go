package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func RouterLog(log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.WithFields(logrus.Fields{
			"service":     "ping",
			"http_method": ctx.Method(),
			"ip_address":  ctx.IP(),
			"hostname":    ctx.Hostname(),
			"method":      fmt.Sprintf(`[%s] %s`, ctx.Method(), ctx.OriginalURL()),
			"uri":         ctx.OriginalURL(),
			"body":        ctx.Body(),
			"category":    "router",
		}).Info("OK")
		return ctx.Next()
	}
}
