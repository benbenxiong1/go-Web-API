package models

type UserInfo struct {
	Model
	Name string `json:"name"`
	Email string `json:"email"`
}
