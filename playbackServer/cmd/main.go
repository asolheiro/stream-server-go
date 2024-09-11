package main

import (
	"github.com/labstack/echo"
	"github.com/rmndvngrpslhr/stream-server-go/playbackServer/service"
	"net/http"
)

func main() {
	server := echo.New()

	server.GET("/healthcheck", HealthCheck)
	server.GET("/live/:live/*", service.ServeStream())

	server.Logger.Fatal(server.Start(":8001"))
}

// HealthCheck defines a simple answer to a HealthCheck route
func HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Playback App - STATUS: HEALTHY")
}
