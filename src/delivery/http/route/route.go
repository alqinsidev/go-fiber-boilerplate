package route

import (
	"alqinsidev.io/go-clean-architecture-fiber/src/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	Fiber *fiber.App
	H     *http.Handler
	RL    fiber.Handler
}

func (r *RouteConfig) Setup() {
	r.SetupPingRoute()
}

func (r *RouteConfig) SetupPingRoute() {
	r.Fiber.Use(r.RL)
	r.Fiber.Get("/status", r.H.Status)
}
