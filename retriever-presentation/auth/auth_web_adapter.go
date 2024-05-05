package auth_presentation

import (
	"github.com/gofiber/fiber/v2"
	"retriever-core/auth/dto/request"
	"retriever-core/auth/primary"
)

type AuthWebAdapter struct {
	authPort auth_primary.AuthPrimaryPort
}

func NewAuthWebAdapter(authPort auth_primary.AuthPrimaryPort) *AuthWebAdapter {
	return &AuthWebAdapter{
		authPort: authPort,
	}
}

func (r *AuthWebAdapter) SignIn(ctx *fiber.Ctx) error {
	dto := new(auth_request.SignInReqDTO)
	_ = ctx.BodyParser(dto)
	err := r.authPort.SignIn(*dto)
	if err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.SendStatus(200)
}
