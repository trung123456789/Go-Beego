package models

type User struct {
	Id   string `json:"user_id"`
	Name string `json:"user_name"`
	Age  int    `json:"age"`
}
