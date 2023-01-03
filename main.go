package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"log"
	"shortened_link/handler"
	"shortened_link/repository"
	"shortened_link/service"
)

func main() {
	db := repository.Initialize()
	r := repository.UserRepositoryImpl{
		PostgresDb: db,
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error to load env file!!")
	}
	h := handler.NewUserHandler(&r)

	urlService := service.UrlServiceImpl{}
	urlHandler := handler.NewUrlHandler(&urlService)

	e := echo.New()
	e.POST("/shorted", urlHandler.CreateShortedUrl)
	e.GET("/:shortedUrl", urlHandler.GetUrlFromShortedUrl)
	e.POST("/login", h.Login)
	e.POST("/signup", h.SignUp)
	e.Logger.Fatal(e.Start(":3000"))
}
