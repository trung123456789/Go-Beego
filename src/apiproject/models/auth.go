package models

type LoginRequest struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type LogoutRequest struct {
	AuthCd string `json:"auth_cd"`
}

type LoginResponse struct {
	AuthCd string `json:"auth_cd"`
	UserId string `json:"user_id"`
}
