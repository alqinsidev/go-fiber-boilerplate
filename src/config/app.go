package config

import (
	"alqinsidev.io/go-clean-architecture-fiber/src/delivery/http"
	"alqinsidev.io/go-clean-architecture-fiber/src/delivery/http/middleware"
	"alqinsidev.io/go-clean-architecture-fiber/src/delivery/http/route"
	"alqinsidev.io/go-clean-architecture-fiber/src/repository"
	"alqinsidev.io/go-clean-architecture-fiber/src/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	Fiber *fiber.App
	Log   *logrus.Logger
}

func Bootstrap(config *BootstrapConfig) {
	Repository := repository.Init(config.Log)
	Usecase := usecase.Init(config.Log, Repository)
	Handler := http.Init(config.Log, Usecase)

	routerLogger := middleware.RouterLog(config.Log)
	routeConfig := route.RouteConfig{
		Fiber: config.Fiber,
		H:     Handler,
		RL:    routerLogger,
	}

	routeConfig.Setup()
}
