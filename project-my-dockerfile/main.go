package main

import (
	// "database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		fmt.Println("Ouch :(")
		return c.HTML(http.StatusOK, "Success \n")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "77"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
