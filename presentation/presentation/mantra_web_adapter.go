package presentation

import (
	"application/domain/mantra/dto/request"
	"application/port/primary"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MantraWebAdapter struct {
	mantraPort primary.MantraPrimaryPort
}

func New(mantraPort primary.MantraPrimaryPort) *MantraWebAdapter {
	return &MantraWebAdapter{
		mantraPort: mantraPort,
	}
}

func (r *MantraWebAdapter) ReadAllMantras(ctx *fiber.Ctx) error {
	return ctx.JSON(r.mantraPort.ReadAllMantras())
}

func (r *MantraWebAdapter) CreateMantra(ctx *fiber.Ctx) error {
	dto := new(request.CreateMantraReqDTO)
	_ = ctx.BodyParser(dto)
	err := r.mantraPort.CreateMantra(*dto)
	if err != nil {
		return ctx.SendStatus(500)
	}
	return ctx.SendStatus(201)
}

func (r *MantraWebAdapter) DeleteMantra(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := r.mantraPort.DeleteMantra(uuid.MustParse(id))
	if err != nil {
		return ctx.SendStatus(500)
	}
	return ctx.SendStatus(201)
}
