package use_case

import (
	"application/domain/auth/dto/request"
)

type AuthUseCase struct {
	signInUC SignInUseCase
}

func NewAuthUseCase(signInUseCase SignInUseCase) *AuthUseCase {
	return &AuthUseCase{
		signInUC: signInUseCase,
	}
}

func (r *AuthUseCase) SignIn(dto request.SignInReqDTO) error {
	return r.signInUC.Execute(dto)
}
