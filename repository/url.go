package repository

import "shortened_link/model"

type UrlRepository interface {
	Create(shortUrl *model.ShortedUrl) (*model.ShortedUrl, error)
	Count() int64
}
