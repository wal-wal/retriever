package user_secondary

import (
	"retriever-core/user/model"
)

type UserSecondaryPort interface {
	FindExistByUserId(userId string) bool
	CreateUser(user user_model.User) error
	GetUser(userId string) (user_model.User, error)
}
