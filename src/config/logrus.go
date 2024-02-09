package config

import (
	"github.com/sirupsen/logrus"
)

func InitLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
