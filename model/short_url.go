package model

import (
	"time"
)

type ShortedUrl struct {
	Id int `json:"id"`
	// Redirect    string `json:"redirect"`
	LongUrl    string    `json:"url"`
	ShortedUrl string    `json:"shorted_url"`
	ExpiredAt  time.Time `json:"expired_at"`
	Custom     bool      `json:"custom"`
}

type UrlCreationRequest struct {
	LongUrl   string `json:"url"`
	CustomUrl string `json:"shortUrl"`
}

type UrlUpdateRequest struct {
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

func (s ShortedUrl) IsExpire() bool {
	return s.ExpiredAt.Before(time.Now())
}
