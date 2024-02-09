package repository

import (
	"alqinsidev.io/go-clean-architecture-fiber/src/domain/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	l *logrus.Logger
}

type RepositoryInterface interface {
	PingStatus(ctx *fiber.Ctx) (Ping *entity.Ping, err error)
}

func Init(log *logrus.Logger) RepositoryInterface {
	return &Repository{
		l: log,
	}
}

func (r *Repository) PingStatus(ctx *fiber.Ctx) (status *entity.Ping, err error) {
	return &entity.Ping{
		Status: "SERVICE ARE RUNNING",
		Active: true,
	}, nil
}
