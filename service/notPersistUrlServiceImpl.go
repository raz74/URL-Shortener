package service

import (
	"errors"
	"fmt"
	"log"
	"shortened_link/model"
	"shortened_link/repository"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var shortToSrcMap = map[string]model.ShortedUrl{}

type UrlServiceImpl struct {
}

func init() {
	shortToSrcMap, _ = repository.ReadCSVFile("repository/shorted.csv")
	go func() {
		for {
			time.Sleep(5 * time.Second)
			err := repository.WriteCSVFile(shortToSrcMap, "repository/shorted.csv")
			if err != nil {
				log.Fatalln(err)
				return
			}
		}
	}()
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
	if result.IsExpire() {
		delete(shortToSrcMap, shortUrl)
		return &result, !isFound
	}
	return &result, isFound
}

func (u *UrlServiceImpl) generateShortedUrl() (string, error) {
	lenght := len(alphabet)
	shortUrl := ""
	count := len(shortToSrcMap) + 1
	for count > 0 {
		i := count % lenght
		shortUrl += string(alphabet[i-1])
		count = (count - i) / len(alphabet)
	}
	return shortUrl, nil
}
