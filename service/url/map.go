package url

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"shortened_link/model"
	"time"
)

const (
	Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var shortToSrcMap = map[string]model.ShortedUrl{}

type UrlServiceImpl struct {
}

func (u *UrlServiceImpl) AddUrl(srcUrl string) (*model.ShortedUrl, error) {
	var shortUrl model.ShortedUrl
	newShort, err := u.generateShortedUrl()
	shortUrl.ShortedUrl = newShort
	if err != nil {
		return &shortUrl, err
	}
	expireTime := time.Now().Add(24 * 7 * time.Hour)
	shortUrl = model.ShortedUrl{
		LongUrl:    srcUrl,
		ExpiredAt:  expireTime,
		ShortedUrl: newShort,
		Custom:     false,
	}
	shortToSrcMap[newShort] = shortUrl
	fmt.Println("map:", shortToSrcMap)
	return &shortUrl, nil
}

func (u *UrlServiceImpl) AddCustomUrl(customUrl, srcUrl string) (*model.ShortedUrl, error) {
	var shortUrl model.ShortedUrl
	shortUrl.ShortedUrl = customUrl
	// check custom url is unique!
	_, isFound := shortToSrcMap[customUrl]
	if !isFound {
		expireTime := time.Now().Add(24 * 7 * time.Hour)
		shortUrl = model.ShortedUrl{
			LongUrl:    srcUrl,
			ShortedUrl: customUrl,
			ExpiredAt:  expireTime,
			Custom:     true,
		}
		shortToSrcMap[customUrl] = shortUrl
		fmt.Println("map:", shortToSrcMap)
	} else {
		return &shortUrl, errors.New("this short url is already exist!try another")
	}
	return &shortUrl, nil
}

func (u *UrlServiceImpl) GetUrl(shortUrl string) (*model.ShortedUrl, bool) {
	result, isFound := shortToSrcMap[shortUrl]
	if isFound && result.IsExpire() {
		delete(shortToSrcMap, shortUrl)
		return &result, !isFound
	}
	return &result, isFound
}

func (u *UrlServiceImpl) UpdateLongUrl(shorted, newLong string) (*model.ShortedUrl, error) {
	var shortedUrl model.ShortedUrl
	shortedUrl.LongUrl = newLong
	ex := shortToSrcMap[shorted].ExpiredAt
	shortedUrl = model.ShortedUrl{
		LongUrl:    newLong,
		ShortedUrl: shorted,
		ExpiredAt:  ex,
		Custom:     false,
	}
	shortToSrcMap[shorted] = shortedUrl
	return &shortedUrl, nil
}

func (u *UrlServiceImpl) UpdateShortUrl(key, newShort string) (*model.ShortedUrl, error) {
	var shortedUrl model.ShortedUrl
	ex := shortToSrcMap[key].ExpiredAt
	y := shortToSrcMap[key].LongUrl
	//check new short is unique
	_, isFound := shortToSrcMap[newShort]
	if isFound {
		return nil, echo.ErrForbidden
	}
	delete(shortToSrcMap, key)
	key = newShort
	shortedUrl = model.ShortedUrl{
		LongUrl:    y,
		ShortedUrl: newShort,
		ExpiredAt:  ex,
		Custom:     true,
	}
	shortToSrcMap[newShort] = shortedUrl
	return &shortedUrl, nil
}

func (u *UrlServiceImpl) DeleteUrl(shorted string) error {
	delete(shortToSrcMap, shorted)
	return nil
}

func (u *UrlServiceImpl) generateShortedUrl() (string, error) {
	lenght := len(Alphabet)
	shortUrl := ""
	count := len(shortToSrcMap) + 1
	for count > 0 {
		i := count % lenght
		shortUrl += string(Alphabet[i])
		count = (count - i) / len(Alphabet)
	}
	return shortUrl, nil
}
