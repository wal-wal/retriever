package user_primary

import "retriever-core/user/dto/request"

type UserPrimaryPort interface {
	CreateUser(dto user_request.CreateUserReqDTO) error
}
