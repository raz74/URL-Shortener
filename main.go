package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"log"
	"shortened_link/handler"
	"shortened_link/repository"
	"shortened_link/repository/userRepository"
	"shortened_link/service/url"
)

func main() {
	db := repository.Initialize()
	mongodb := repository.MongoInitialize()
	//r := repository.UserRepositoryImpl{
	//	PostgresDb: db,
	//}
	rMongo := userRepository.NewMongoUserRepositoryImpl(mongodb)
	t := repository.TokenRepositoryImp{
		PostgresDb: db,
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error to load env file!!")
	}

	h := handler.NewUserHandler(rMongo, &t)

	urlService := url.PostgresUrlServiceImpl{DB: db}
	urlHandler := handler.NewUrlHandler(&urlService, &t)

	e := echo.New()
	e.POST("/shorted", urlHandler.CreateShortedUrl)
	e.GET("/:shortedUrl", urlHandler.GetUrlFromShortedUrl)
	e.POST("/login", h.Login)
	e.POST("/signup", h.SignUp)
	e.PUT("/:shortedUrl", urlHandler.UpdateUrl)
	e.DELETE("/:shorted", urlHandler.DeleteShortedUrl)
	e.Logger.Fatal(e.Start(":3000"))
}
