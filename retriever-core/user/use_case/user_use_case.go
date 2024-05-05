package user_use_case

import "retriever-core/user/dto/request"

type UserUseCase struct {
	createUserUC CreateUserUseCase
}

func NewUserUseCase(createUserUseCase CreateUserUseCase) *UserUseCase {
	return &UserUseCase{
		createUserUC: createUserUseCase,
	}
}

func (r *UserUseCase) CreateUser(dto user_request.CreateUserReqDTO) error {
	return r.createUserUC.Execute(dto)
}
