package user_use_case

import (
	"fmt"
	"retriever-core/user/dto/request"
	"retriever-core/user/model"
	"retriever-core/user/secondary"
)

type CreateUserUseCase struct {
	userPort user_secondary.UserSecondaryPort
}

func NewCreateUserUseCase(userPort user_secondary.UserSecondaryPort) *CreateUserUseCase {
	return &CreateUserUseCase{
		userPort: userPort,
	}
}

func (r *CreateUserUseCase) Execute(dto user_request.CreateUserReqDTO) error {
	isDuplicate := r.userPort.FindExistByUserId(dto.UserId)
	if isDuplicate {
		return fmt.Errorf("이미 사용중인 아이디입니다")
	}

	user := user_model.User{UserId: dto.UserId, Password: dto.Password, Name: dto.Name}
	return r.userPort.CreateUser(user)
}
