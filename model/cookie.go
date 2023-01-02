package model

import "time"

type SessionCookie struct {
	Name   string
	Value  string
	Expire time.Time
}

func (s SessionCookie) IsExpire() bool {
	return s.Expire.Before(time.Now())
}
