package models

type LoginRequest struct {
	UserId   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LogoutRequest struct {
	AuthCd string `json:"auth_cd"`
}

type LoginResponse struct {
	AuthCd string `json:"auth_cd"`
	UserId string `json:"user_id"`
}
