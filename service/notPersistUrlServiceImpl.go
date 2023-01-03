package service

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"shortened_link/repository"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var shortToSrcMap = map[string]string{}

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

func (u *UrlServiceImpl) AddUrl(srcUrl string) (string, error) {
	shortUrl, err := u.generateShortedUrl(srcUrl)
	if err != nil {
		return "", err
	}
	shortToSrcMap[shortUrl] = srcUrl
	fmt.Println("map:", shortToSrcMap)
	return shortUrl, nil
}

func (u *UrlServiceImpl) AddCustomUrl(customUrl, srcUrl string) (string, error) {
	shortUrl := customUrl
	// check custom url is unique!
	_, isFound := shortToSrcMap[shortUrl]
	if !isFound {
		shortToSrcMap[shortUrl] = srcUrl
		fmt.Println("map:", shortToSrcMap)
	} else {
		return "This short Url is already exist!", echo.ErrForbidden
	}
	return shortUrl, nil
}

func (u *UrlServiceImpl) GetUrl(shortUrl string) (string, bool) {
	dest, isFound := shortToSrcMap[shortUrl]
	return dest, isFound
}

func (u *UrlServiceImpl) generateShortedUrl(url string) (string, error) {
	lenght := len(alphabet)
	shortUrl := ""
	count := len(shortToSrcMap) + 1
	for count > 0 {
		i := count % lenght
		shortUrl += string(alphabet[i-1])
		count = (count - i) / len(alphabet)
	}
	shortToSrcMap[shortUrl] = url
	return shortUrl, nil
}
