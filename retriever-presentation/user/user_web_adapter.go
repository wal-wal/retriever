package user_presentation

import (
	"errors"
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
	if err := createUserReqDTOValidator(*dto); err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	if err := r.userPort.CreateUser(*dto); err != nil {
		return ctx.Status(409).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.SendStatus(201)
}

func createUserReqDTOValidator(dto user_request.CreateUserReqDTO) error {
	if dto.UserId == "" {
		return errors.New("아이디를 입력해주세요")
	}
	if dto.Password == "" {
		return errors.New("비밀번호를 입력해주세요")
	}
	if dto.Name == "" {
		return errors.New("이름을 입랙해주세요")
	}
	return nil
}
