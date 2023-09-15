package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.HideBanner = true
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{"msg": "hello world"})
	})

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err.Error())

	}
}
