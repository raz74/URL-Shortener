package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"log"
	"shortened_link/handler"
	"shortened_link/repository"
	"time"
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

	//to do : call read csv file to load before requests start
	// Read shorted urls from 'shorted.csv'
	handler.MyMap, _ = repository.ReadCSVFile("repository/shorted.csv")
	go func() {
		for {
			time.Sleep(5 * time.Second)
			err := repository.WriteCSVFile(handler.MyMap, "repository/shorted.csv")
			if err != nil {
				log.Fatalln(err)
				return
			}
		}
	}()

	e := echo.New()
	e.POST("/shorted", handler.CreateShortedUrl)
	e.GET("/:shortedUrl", handler.GetUrlFromShortedUrl)
	e.POST("/login", h.Login)
	e.POST("/signup", h.SignUp)
	e.Logger.Fatal(e.Start(":3000"))
}
