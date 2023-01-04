package model

import "time"

type ShortedUrl struct {
	Id int `json:"id"`
	// Redirect    string `json:"redirect"`
	LongUrl    string    `json:"url"`
	ShortedUrl string    `json:"shorted_url"`
	ExpiredAt  time.Time `json:"expired_at"`
}

type UrlCreationRequest struct {
	LongUrl   string `json:"url"`
	CustomUrl string `json:"customUrl"`

	//Random bool `json:"random"`
}

func (s ShortedUrl) IsExpire() bool {
	return s.ExpiredAt.Before(time.Now())
}
