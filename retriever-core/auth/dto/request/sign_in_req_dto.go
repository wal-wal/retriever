package auth_request

type SignInReqDTO struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}
