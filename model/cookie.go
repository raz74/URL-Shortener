package model

import "time"

type SessionCookie struct {
	UserID int `gorm:"primaryKey"`
	Value  string
	Expire time.Time
}
