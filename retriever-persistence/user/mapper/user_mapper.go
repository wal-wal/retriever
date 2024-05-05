package persistence_user

import (
	core_user "retriever-core/user/model"
	"retriever-persistence/user/entity"
)

type UserMapper struct {
}

func (r *UserMapper) ToDomain(entity persistence_user.UserEntity) core_user.User {
	return core_user.User{
		UserId:   entity.UserId,
		Password: entity.Password,
		Name:     entity.Name,
	}
}

func (r *UserMapper) ToEntity(model core_user.User) persistence_user.UserEntity {
	return persistence_user.UserEntity{
		UserId:   model.UserId,
		Password: model.Password,
		Name:     model.Name,
	}
}
