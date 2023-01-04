package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"shortened_link/model"
	"shortened_link/service"
)

type UrlHandler struct {
	urlService service.UrlService
}

func NewUrlHandler(urlService service.UrlService) *UrlHandler {
	return &UrlHandler{urlService: urlService}
}

func (u *UrlHandler) CreateShortedUrl(c echo.Context) error {
	header := c.Request().Header.Get("Authorization") // Token sadasdfasdfsa
	err := CheckHeaderAuthorize(header)
	if err != nil {
		return err
	}

	var request *model.UrlCreationRequest
	var shortUrl *model.ShortedUrl
	if err := c.Bind(&request); err != nil {
		return echo.ErrBadRequest
	}
	if len(request.CustomUrl) > 0 {
		shortUrl, err = u.urlService.AddCustomUrl(request.CustomUrl, request.LongUrl)
		if err != nil {
			return echo.ErrForbidden
		}
	} else {
		shortUrl, err = u.urlService.AddUrl(request.LongUrl)

		if err != nil {
			return echo.ErrInternalServerError
		}
	}

	return c.JSON(http.StatusOK, shortUrl)
}

func (u *UrlHandler) GetUrlFromShortedUrl(c echo.Context) error {
	shortedUrl := c.Param("shortedUrl")

	url, found := u.urlService.GetUrl(shortedUrl)

	if !found {
		return c.JSON(http.StatusNotFound, "This shorted_url is not existing!")
	}

	return c.Redirect(http.StatusTemporaryRedirect, url.LongUrl)
}
