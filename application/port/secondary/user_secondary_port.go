package secondary

import (
	"application/domain/user/model"
)

type UserSecondaryPort interface {
	FindExistByUserId(userId string) bool
	CreateUser(user model.User) error
	GetUser(userId string) (model.User, error)
}
