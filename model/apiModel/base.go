package apimodel

type LoginParam struct {
	Username string `json:"username" example:"my_username"`
	Password string `json:"password" example:"my_password"`
}
