package auth_primary

import "retriever-core/auth/dto/request"

type AuthPrimaryPort interface {
	SignIn(dto auth_request.SignInReqDTO) error
}
