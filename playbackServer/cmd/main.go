package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	server := echo.New()

	server.GET(`/healthcheck`, HealthCheck)

	server.Logger.Fatal(server.Start(":8001"))
}

// HealthCheck defines a simple answer to a HealthCheck route
func HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Playback App - STATUS: HEALTHY")
}
