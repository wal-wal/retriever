package use_case

import (
	"application/domain/auth/dto/request"
	"application/port/secondary"
	"errors"
)

type SignInUseCase struct {
	userPort secondary.UserSecondaryPort
}

func NewSignInUseCase(userPort secondary.UserSecondaryPort) *SignInUseCase {
	return &SignInUseCase{
		userPort: userPort,
	}
}

func (r *SignInUseCase) Execute(dto request.SignInReqDTO) error {
	user, err := r.userPort.GetUser(dto.UserId)
	if err != nil {
		return errors.New("존재하지 않는 아이디입니다.")
	}
	if user.Password != dto.Password {
		return errors.New("비밀번호가 일치하지 않습니다.")
	}
	return nil
}
