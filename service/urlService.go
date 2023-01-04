package service

import "shortened_link/model"

type UrlService interface {
	AddUrl(string) (*model.ShortedUrl, error)
	AddCustomUrl(string, string) (*model.ShortedUrl, error)
	GetUrl(string) (*model.ShortedUrl, bool)
}
