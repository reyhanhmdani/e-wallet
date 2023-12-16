package dto

type AuthReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
