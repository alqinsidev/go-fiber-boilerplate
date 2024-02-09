package http

import (
	"alqinsidev.io/go-clean-architecture-fiber/src/usecase"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log *logrus.Logger
	uc  usecase.UsecaseInterface
}

func Init(log *logrus.Logger, usecase usecase.UsecaseInterface) *Handler {
	return &Handler{
		log: log,
		uc:  usecase,
	}
}
