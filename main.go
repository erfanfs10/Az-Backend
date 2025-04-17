package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	fmt.Println("Running ...")

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Az-Lab-Backend")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
