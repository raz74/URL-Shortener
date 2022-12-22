package main

import (
	"github.com/labstack/echo"
	"log"
	"shortened_link/handler"
	"shortened_link/repository"
	"time"
)

func main() {
	//to do : call read csv file to load before requests start
	// Read shorted urls from 'shorted.csv'
	handler.MyMap, _ = repository.ReadCSVFile("repository/shorted.csv")
	go func() {
		for {
			time.Sleep(5 * time.Second)
			if handler.CountDif > 5 {
				err := repository.WriteCSVFile(handler.MyMap, "repository/shorted.csv")
				if err != nil {
					log.Fatalln(err)
					return
				}
				handler.CountDif = 0
			}
		}

	}()

	e := echo.New()
	e.POST("/shorted", handler.CreateShortedUrl)
	e.GET("/:shortedUrl", handler.GetUrlFromShortedUrl)
	e.Logger.Fatal(e.Start(":3000"))

	// to do: call write in csv at the end of main!

}
