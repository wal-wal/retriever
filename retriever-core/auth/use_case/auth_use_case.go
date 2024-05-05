package auth_use_case

import auth_request "retriever-core/auth/dto/request"

type AuthUseCase struct {
	signInUC SignInUseCase
}

func NewAuthUseCase(signInUseCase SignInUseCase) *AuthUseCase {
	return &AuthUseCase{
		signInUC: signInUseCase,
	}
}

func (r *AuthUseCase) SignIn(dto auth_request.SignInReqDTO) error {
	return r.signInUC.Execute(dto)
}
