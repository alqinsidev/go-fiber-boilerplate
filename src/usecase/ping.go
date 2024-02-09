package usecase

import (
	"alqinsidev.io/go-clean-architecture-fiber/src/domain/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func (u *Usecase) PingStatus(ctx *fiber.Ctx) (status *model.PingStatusResponse, err error) {
	data, err := u.r.PingStatus(ctx)
	if err != nil {
		return
	}

	status = new(model.PingStatusResponse)
	err = copier.Copy(&status, &data)
	if err != nil {
		return
	}
	return
}
