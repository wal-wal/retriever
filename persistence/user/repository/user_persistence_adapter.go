package repository

import (
	"application/domain/user/model"
	"database/sql"
	"persistence/user/entity"
	"persistence/user/mapper"
)

type UserPersistenceAdapter struct {
	db     *sql.DB
	mapper mapper.UserMapper
}

func NewUserPersistenceAdapter(db *sql.DB) *UserPersistenceAdapter {
	return &UserPersistenceAdapter{
		db: db,
	}
}

func (r *UserPersistenceAdapter) FindExistByUserId(userId string) bool {
	userEntity := new(entity.UserEntity)
	err := r.db.QueryRow("SELECT * FROM tbl_user WHERE userId = ?", userId).Scan(&userEntity.UserId, &userEntity.Password, &userEntity.Name)
	return userEntity.UserId != "" && err == nil
}

func (r *UserPersistenceAdapter) CreateUser(user model.User) error {
	_, err := r.db.Exec("INSERT INTO tbl_user VALUES (?, ?, ?)", user.UserId, user.Password, user.Name)
	return err
}

func (r *UserPersistenceAdapter) GetUser(userId string) (model.User, error) {
	userEntity := new(entity.UserEntity)
	err := r.db.QueryRow("SELECT * FROM tbl_user WHERE userId = ?", userId).Scan(&userEntity.UserId, &userEntity.Password, &userEntity.Name)
	return r.mapper.ToDomain(*userEntity), err
}
