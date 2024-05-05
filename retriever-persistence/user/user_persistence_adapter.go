package user_persistence

import (
	"retriever-core/user/model"
	"retriever-persistence/user/mapper"
	"retriever-persistence/user/repository"
)

type UserPersistenceAdapter struct {
	repository user_repository.UserRepository
	mapper     user_mapper.UserMapper
}

func NewUserPersistenceAdapter(repository user_repository.UserRepository) *UserPersistenceAdapter {
	return &UserPersistenceAdapter{
		repository: repository,
	}
}

func (r *UserPersistenceAdapter) FindExistByUserId(userId string) bool {
	return r.repository.FindExistByUserId(userId)
}

func (r *UserPersistenceAdapter) CreateUser(user user_model.User) error {
	return r.repository.CreateUser(r.mapper.ToEntity(user))
}

func (r *UserPersistenceAdapter) GetUser(userId string) (user_model.User, error) {
	userEntity, err := r.repository.GetUser(userId)
	return r.mapper.ToDomain(userEntity), err
}
