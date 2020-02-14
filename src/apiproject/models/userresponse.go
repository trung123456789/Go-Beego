package models

type UserResponse struct {
	RecordNum int64      `json:"record_num"`
	UserList  []UserInfo `json:"user_list"`
}
