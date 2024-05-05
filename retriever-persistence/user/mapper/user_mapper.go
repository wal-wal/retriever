package user_mapper

import (
	"retriever-core/user/model"
	"retriever-persistence/user/entity"
)

type UserMapper struct {
}

func (r *UserMapper) ToDomain(entity user_entity.UserEntity) user_model.User {
	return user_model.User{
		UserId:   entity.UserId,
		Password: entity.Password,
		Name:     entity.Name,
	}
}

func (r *UserMapper) ToEntity(model user_model.User) user_entity.UserEntity {
	return user_entity.UserEntity{
		UserId:   model.UserId,
		Password: model.Password,
		Name:     model.Name,
	}
}
