package main

import (
	"fmt"
	"github.com/labstack/echo"
	"io"
	"log"
	"net/http"
)

func main() {
	server := echo.New()

	server.GET("/healthcheck", func(context echo.Context) error {
		return context.String(http.StatusOK, "OK")
	})

	server.POST("/auth", func(context echo.Context) error {
		log.Default().Println("Running Auth...")
		body := context.Request().Body
		defer body.Close()

		fields, _ := io.ReadAll(body)
		fmt.Println(string(fields))

		return context.String(http.StatusOK, "WORKING")
	})

	server.Logger.Fatal(
		server.Start(
			":8000",
		),
	)
}
