package user_repository

import (
	"database/sql"
	"retriever-persistence/user/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindExistByUserId(userId string) bool {
	userEntity := new(user_entity.UserEntity)
	err := r.db.QueryRow("SELECT * FROM tbl_user WHERE userId = ?", userId).Scan(&userEntity.UserId, &userEntity.Password, &userEntity.Name)
	return userEntity.UserId != "" && err == nil
}

func (r *UserRepository) CreateUser(entity user_entity.UserEntity) error {
	_, err := r.db.Exec("INSERT INTO tbl_user VALUES (?, ?, ?)", entity.UserId, entity.Password, entity.Name)
	return err
}

func (r *UserRepository) GetUser(userId string) (entity user_entity.UserEntity, err error) {
	err = r.db.QueryRow("SELECT * FROM tbl_user WHERE userId = ?", userId).Scan(&entity.UserId, &entity.Password, &entity.Name)
	return entity, err
}
