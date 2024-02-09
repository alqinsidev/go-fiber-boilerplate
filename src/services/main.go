package main

import (
	"fmt"

	"alqinsidev.io/go-clean-architecture-fiber/src/config"
	"github.com/spf13/viper"
)

func main() {
	log := config.InitLogger()
	fiber := config.InitFiber()
	_ = config.InitViper()

	config.Bootstrap(&config.BootstrapConfig{
		Fiber: fiber,
		Log:   log,
	})

	appPort := viper.GetInt("APP_PORT")
	err := fiber.Listen(fmt.Sprintf(":%d", appPort))
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
}
