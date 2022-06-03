package main

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func main() {
	server := echo.New()
	server.HideBanner = true
	server.GET("/health", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "OK")
	})
	log.Fatal(server.Start(":8080"))
}
