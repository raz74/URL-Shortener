package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"shortened_link/model"
	"strconv"
)

var MyMap = make(map[string]string)

func CreateShortedUrl(c echo.Context) error {
	var request *model.ShortendUrl
	if err := c.Bind(&request); err != nil {
		return echo.ErrBadRequest
	}
	//shortUrl := GenerateShortLink(model.UrlCreationRequest{})
	shortUrl, err := GenerateShortedUrl(request.Url)
	if err != nil {
		return err
	}
	MyMap[request.Redirect] = request.Url
	//fmt.Println("map:", MyMap)
	return c.JSON(http.StatusOK, shortUrl)
}

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateShortedUrl(url string) (string, error) {
	lenght := len(alphabet)
	MyMap[strconv.Itoa(lenght)] = url
	shortUrl := ""
	count := lenght + 1
	for count > 0 {
		i := count % lenght
		shortUrl += string(alphabet[i])
		count = (count - i) / len(alphabet)
	}
	MyMap[shortUrl] = url
	return shortUrl, nil
}
