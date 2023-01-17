package model

type User struct {
	Id       int             `json:"id"`
	Name     string          `json:"name"`
	Password string          `json:"password"`
	Email    string          `json:"email" gorm:"unique"`
	Cookie   []SessionCookie `gorm:"constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	Custom   bool            `json:"custom"`
}
