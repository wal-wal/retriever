package auth_presentation

import (
	"errors"
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
	if err := signInReqDTOValidator(*dto); err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	if err := r.authPort.SignIn(*dto); err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.SendStatus(200)
}

func signInReqDTOValidator(dto auth_request.SignInReqDTO) error {
	if dto.UserId == "" {
		return errors.New("아이디를 입력해주세요")
	}
	if dto.Password == "" {
		return errors.New("비밀번호를 입력해주세요")
	}
	return nil
}
