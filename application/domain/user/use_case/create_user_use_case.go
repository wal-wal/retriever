package use_case

import (
	"application/domain/user/dto/request"
	"application/domain/user/model"
	"application/port/secondary"
	"fmt"
)

type CreateUserUseCase struct {
	userPort secondary.UserSecondaryPort
}

func NewCreateUserUseCase(userPort secondary.UserSecondaryPort) *CreateUserUseCase {
	return &CreateUserUseCase{
		userPort: userPort,
	}
}

func (r *CreateUserUseCase) Execute(dto request.CreateUserReqDTO) error {
	isDuplicate := r.userPort.FindExistByUserId(dto.UserId)
	if isDuplicate {
		return fmt.Errorf("이미 사용중인 아이디입니다")
	}
	user := model.User{UserId: dto.UserId, Password: dto.Password, Name: dto.Name}
	return r.userPort.CreateUser(user)
}
