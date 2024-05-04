package primary

import (
	"application/domain/auth/dto/request"
)

type AuthPrimaryPort interface {
	SignIn(dto request.SignInReqDTO) error
}
