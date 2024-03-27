package presentation

import (
	request2 "application/domain/user/dto/request"
	"application/port/primary"
	"github.com/gofiber/fiber/v2"
)

type UserWebAdapter struct {
	userPort primary.UserPrimaryPort
}

func NewUserWebAdapter(userPort primary.UserPrimaryPort) *UserWebAdapter {
	return &UserWebAdapter{
		userPort: userPort,
	}
}

func (r *UserWebAdapter) CreateUser(ctx *fiber.Ctx) error {
	dto := new(request2.CreateUserReqDTO)
	_ = ctx.BodyParser(dto)
	err := r.userPort.CreateUser(*dto)
	if err != nil {
		return ctx.Status(409).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.SendStatus(201)
}
