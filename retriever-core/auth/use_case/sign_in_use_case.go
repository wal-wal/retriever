package auth_use_case

import (
	"errors"
	"retriever-core/auth/dto/request"
	"retriever-core/user/secondary"
)

type SignInUseCase struct {
	userPort user_secondary.UserSecondaryPort
}

func NewSignInUseCase(userPort user_secondary.UserSecondaryPort) SignInUseCase {
	return SignInUseCase{
		userPort: userPort,
	}
}

func (r *SignInUseCase) Execute(dto auth_request.SignInReqDTO) error {
	user, err := r.userPort.GetUser(dto.UserId)
	if err != nil {
		return errors.New("존재하지 않는 아이디입니다.")
	}
	if user.Password != dto.Password {
		return errors.New("비밀번호가 일치하지 않습니다.")
	}
	return nil
}
