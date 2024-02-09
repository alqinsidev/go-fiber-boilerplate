package usecase

import (
	"alqinsidev.io/go-clean-architecture-fiber/src/domain/model"
	"alqinsidev.io/go-clean-architecture-fiber/src/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	l *logrus.Logger
	r repository.RepositoryInterface
}

type UsecaseInterface interface {
	PingStatus(ctx *fiber.Ctx) (status *model.PingStatusResponse, err error)
}

func Init(log *logrus.Logger, repository repository.RepositoryInterface) UsecaseInterface {
	return &Usecase{
		l: log,
		r: repository,
	}
}
