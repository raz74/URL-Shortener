package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"shortened_link/model"
)

var MyMap = make(map[string]string)

// count how many diffrent item have to write
var CountDif = 0

func CreateShortedUrl(c echo.Context) error {
	var request *model.UrlCreationRequest
	if err := c.Bind(&request); err != nil {
		return echo.ErrBadRequest
	}
	shortUrl, err := GenerateShortedUrl(request.LongUrl)
	if err != nil {
		return err
	}
	MyMap[shortUrl] = request.LongUrl
	fmt.Println("map:", MyMap)
	CountDif++
	return c.JSON(http.StatusOK, shortUrl)
}

func GetUrlFromShortedUrl(c echo.Context) error {
	shortedUrl := c.Param("shortedUrl")

	url, found := MyMap[shortedUrl]
	if !found {
		return c.JSON(http.StatusNotFound, "This shorted_url is not existing!")
	}
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateShortedUrl(url string) (string, error) {
	lenght := len(alphabet)
	shortUrl := ""
	count := len(MyMap) + 1
	for count > 0 {
		i := count % lenght
		shortUrl += string(alphabet[i-1])
		count = (count - i) / len(alphabet)
	}
	MyMap[shortUrl] = url
	return shortUrl, nil
}
