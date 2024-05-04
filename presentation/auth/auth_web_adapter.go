package auth

import (
	"application/domain/auth/dto/request"
	"application/port/primary"
	"github.com/gofiber/fiber/v2"
)

type AuthWebAdapter struct {
	authPort primary.AuthPrimaryPort
}

func NewAuthWebAdapter(authPort primary.AuthPrimaryPort) *AuthWebAdapter {
	return &AuthWebAdapter{
		authPort: authPort,
	}
}

func (r *AuthWebAdapter) SignIn(ctx *fiber.Ctx) error {
	dto := new(request.SignInReqDTO)
	_ = ctx.BodyParser(dto)
	err := r.authPort.SignIn(*dto)
	if err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.SendStatus(200)
}
