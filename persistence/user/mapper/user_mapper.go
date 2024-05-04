package mapper

import (
	"application/domain/user/model"
	"persistence/user/entity"
)

type UserMapper struct {
}

func (r *UserMapper) ToDomain(entity entity.UserEntity) model.User {
	return model.User{
		UserId:   entity.UserId,
		Password: entity.Password,
		Name:     entity.Name,
	}
}

func (r *UserMapper) ToEntity(model model.User) entity.UserEntity {
	return entity.UserEntity{
		UserId:   model.UserId,
		Password: model.Password,
		Name:     model.Name,
	}
}
