package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Cookie   string `json:"cookie"`
}
