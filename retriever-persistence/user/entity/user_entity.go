package persistence_user

type UserEntity struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
