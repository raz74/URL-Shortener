package url

import "shortened_link/model"

type UrlService interface {
	AddUrl(string) (*model.ShortedUrl, error)
	AddCustomUrl(string, string) (*model.ShortedUrl, error)
	GetUrl(string) (*model.ShortedUrl, bool)
	UpdateShortUrl(string, string) (*model.ShortedUrl, error)
	UpdateLongUrl(string, string) (*model.ShortedUrl, error)
	DeleteUrl(shorted string) error
}
