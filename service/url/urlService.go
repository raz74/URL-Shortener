package url

import (
	"shortened_link/model"
	"shortened_link/repository"
	"time"
)

type UrlService interface {
	AddUrl(string) (*model.ShortedUrl, error)
	AddCustomUrl(string, string) (*model.ShortedUrl, error)
	GetUrl(string) (*model.ShortedUrl, bool)
	UpdateShortUrl(string, string) (*model.ShortedUrl, error)
	UpdateLongUrl(string, string) (*model.ShortedUrl, error)
	DeleteUrl(shorted string) error
}

type urlServiceImpl struct {
	repo repository.UrlRepository
}

func NewUrlService(repo repository.UrlRepository) *urlServiceImpl {
	return &urlServiceImpl{repo: repo}
}

func (u *urlServiceImpl) AddUrl(srcUrl string) (*model.ShortedUrl, error) {
	count := u.repo.Count()
	newShort := u.generateShortedUrl(count)
	var shortUrl model.ShortedUrl
	shortUrl.ShortedUrl = newShort

	expireTime := time.Now().Add(24 * 7 * time.Hour)
	shortUrl = model.ShortedUrl{
		LongUrl:    srcUrl,
		ShortedUrl: newShort,
		ExpiredAt:  expireTime,
		Custom:     false,
	}

	return u.repo.Create(&shortUrl)
}

func (u *urlServiceImpl) AddCustomUrl(s string, s2 string) (*model.ShortedUrl, error) {
	//TODO implement me
	panic("implement me")
}

func (u *urlServiceImpl) GetUrl(s string) (*model.ShortedUrl, bool) {
	//TODO implement me
	panic("implement me")
}

func (u *urlServiceImpl) UpdateShortUrl(s string, s2 string) (*model.ShortedUrl, error) {
	//TODO implement me
	panic("implement me")
}

func (u *urlServiceImpl) UpdateLongUrl(s string, s2 string) (*model.ShortedUrl, error) {
	//TODO implement me
	panic("implement me")
}

func (u *urlServiceImpl) DeleteUrl(shorted string) error {
	//TODO implement me
	panic("implement me")
}

func (u *urlServiceImpl) generateShortedUrl(count int64) string {
	lenght := int64(len(Alphabet))
	shortUrl := ""
	for count > 0 {
		i := count % lenght
		shortUrl += string(Alphabet[i])
		count = (count - i) / lenght
	}
	return shortUrl

}
