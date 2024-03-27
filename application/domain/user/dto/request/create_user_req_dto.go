package request

type CreateUserReqDTO struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
