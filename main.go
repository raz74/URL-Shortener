package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"log"
	"shortened_link/handler"
	"shortened_link/repository"
	"shortened_link/repository/token"
	urlRepository "shortened_link/repository/url"
	"shortened_link/repository/userRepository"
	"shortened_link/service/url"
)

func main() {
	//db := repository.Initialize()
	mongodb := repository.MongoInitialize()
	//r := repository.PostgresUserRepositoryImpl{
	//	PostgresDb: db,
	//}
	rMongo := userRepository.NewMongoUserRepoImpl(mongodb)

	//t := token.PostTokenRepoImp{
	//	PostgresDb: db,
	//}
	tMongo := token.NewMongoTokenRepoImp(mongodb)

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error to load env file!!")
	}

	h := handler.NewUserHandler(rMongo, tMongo)

	//urlService := url.PostgresUrlServiceImpl{DB: db}

	urlMongo := urlRepository.NewMongoUrlRepoImp(mongodb)
	urlService := url.NewUrlService(urlMongo)
	urlHandler := handler.NewUrlHandler(urlService, tMongo)

	e := echo.New()
	e.POST("/shorted", urlHandler.CreateShortedUrl)
	e.GET("/:shortedUrl", urlHandler.GetUrlFromShortedUrl)
	e.POST("/login", h.Login)
	e.POST("/signup", h.SignUp)
	e.PUT("/:shortedUrl", urlHandler.UpdateUrl)
	e.DELETE("/:shorted", urlHandler.DeleteShortedUrl)
	e.Logger.Fatal(e.Start(":3000"))
}
