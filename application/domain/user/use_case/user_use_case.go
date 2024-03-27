package use_case

import "application/domain/user/dto/request"

type UserUseCase struct {
	createUserUC CreateUserUseCase
}

func NewUserUseCase(createUserUseCase CreateUserUseCase) *UserUseCase {
	return &UserUseCase{
		createUserUC: createUserUseCase,
	}
}

func (r *UserUseCase) CreateUser(dto request.CreateUserReqDTO) error {
	return r.createUserUC.Execute(dto)
}
