package main

import (
	"github.com/labstack/echo/v4"

	"fiveloved/handlers"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", handlers.Index)

	e.Logger.Fatal(e.Start(":3000"))
}
