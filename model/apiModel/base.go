package apimodel

type LoginParam struct {
	Username string `json:"username" example:"my_username"`
	Password string `json:"password" example:"my_password"`
}

type LoginResData struct {
	Token  string `json:"token" example:"eyJhbGciOiJIUzI1NiIkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsiOjE1ODU2nVpZCI6IVzZXJuYW1lIjoic3RyaW5nIn0.HbrhJbblrWLVqle6TI19bGX78ki4x5x1Wxs"`
	Expire string `json:"expire" example:"2020-04-01T12:08:36+08:00"`
}

type UserRegistParam struct {
	Username string `json:"username" example:"my_username"`
	Password string `json:"password" example:"my_password"`
	Email    string `json:"email" example:"my_email@example.com"`
	Nickname string `json:"nickname" example:"my_nickname"`
}
