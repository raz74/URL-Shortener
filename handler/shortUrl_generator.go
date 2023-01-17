package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"shortened_link/model"
	"shortened_link/service/url"
)

type UrlHandler struct {
	urlService url.UrlService
}

func NewUrlHandler(urlService url.UrlService) *UrlHandler {
	return &UrlHandler{urlService: urlService}
}

func (u *UrlHandler) CreateShortedUrl(c echo.Context) error {
	header := c.Request().Header.Get("Authorization") // Token sadasdfasdfsa
	err := CheckHeaderAuthorize(header)
	//if err != nil {
	//	return err
	//}

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

	getUrl, found := u.urlService.GetUrl(shortedUrl)

	if !found {
		return c.JSON(http.StatusNotFound, "This shorted_url is not existing!")
	}
	fmt.Print(getUrl)
	return c.Redirect(http.StatusTemporaryRedirect, getUrl.LongUrl)
}

func (u *UrlHandler) UpdateUrl(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	err := CheckHeaderAuthorize(header)
	//if err != nil {
	//	return err
	//}

	shortedUrl := c.Param("shortedUrl")
	theShort, found := u.urlService.GetUrl(shortedUrl)
	if !found {
		return c.JSON(http.StatusNotFound, "This shorted_url is not existing!")
	}
	custom := theShort.Custom
	var request model.UrlUpdateRequest
	var result *model.ShortedUrl
	if err := c.Bind(&request); err != nil {
		return echo.ErrBadRequest
	}
	//non custom users just can update long url
	if !custom {
		result, err = u.urlService.UpdateLongUrl(shortedUrl, request.LongUrl)
		if err != nil {
			return echo.ErrBadRequest
		}
	}
	if custom {
		result, err = u.urlService.UpdateShortUrl(shortedUrl, request.ShortUrl)
		if err != nil {
			return echo.ErrBadRequest
		}
	}

	return c.JSON(http.StatusOK, result)
}

func (u *UrlHandler) DeleteShortedUrl(c echo.Context) error {
	shorted := c.Param("shortedUrl")

	_, found := u.urlService.GetUrl(shorted)
	if !found {
		return c.JSON(http.StatusNotFound, "This shorted_url is not existing!")
	}

	err := u.urlService.DeleteUrl(shorted)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, "delete successfully")
}
