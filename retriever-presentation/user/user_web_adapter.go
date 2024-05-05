package user_presentation

import (
	"github.com/gofiber/fiber/v2"
	"retriever-core/user/dto/request"
	"retriever-core/user/primary"
)

type UserWebAdapter struct {
	userPort user_primary.UserPrimaryPort
}

func NewUserWebAdapter(userPort user_primary.UserPrimaryPort) *UserWebAdapter {
	return &UserWebAdapter{
		userPort: userPort,
	}
}

func (r *UserWebAdapter) CreateUser(ctx *fiber.Ctx) error {
	dto := new(user_request.CreateUserReqDTO)
	_ = ctx.BodyParser(dto)
	err := r.userPort.CreateUser(*dto)
	if err != nil {
		return ctx.Status(409).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.SendStatus(201)
}
