package primary

import "application/domain/user/dto/request"

type UserPrimaryPort interface {
	CreateUser(dto request.CreateUserReqDTO) error
}
