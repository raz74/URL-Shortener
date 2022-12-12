package main

import (
	"github.com/labstack/echo"
	"shortened_link/handler"
)

func main() {
	e := echo.New()
	e.POST("/shorted", handler.CreateShortedUrl)
	e.Logger.Fatal(e.Start(":3000"))
}
