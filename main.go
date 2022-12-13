package main

import (
	"shortened_link/handler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/shorted", handler.CreateShortedUrl)
	e.GET("/:shortedUrl", handler.GetUrlFromShortedUrl)
	e.Logger.Fatal(e.Start(":3000"))
}
