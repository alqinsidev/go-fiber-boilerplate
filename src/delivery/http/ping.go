package http

import (
	"alqinsidev.io/go-clean-architecture-fiber/src/domain/model"
	"alqinsidev.io/go-clean-architecture-fiber/src/helpers"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Status(ctx *fiber.Ctx) (err error) {
	log := h.log.WithFields(helpers.SetLoggerField(ctx, "usecase", nil))

	data, err := h.uc.PingStatus(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	response := model.HttpResponse[*model.PingStatusResponse]{
		Data:    data,
		Status:  true,
		Message: "Success",
		Code:    2000100,
	}
	err = ctx.JSON(response)
	return
}
