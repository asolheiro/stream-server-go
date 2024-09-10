package main

import (
	"github.com/labstack/echo"
	"github.com/rmndvngrpslhr/stream-server-go/authServer/config/db"
	"github.com/rmndvngrpslhr/stream-server-go/authServer/internal/handler"
	"github.com/rmndvngrpslhr/stream-server-go/authServer/internal/repository"
	"github.com/rmndvngrpslhr/stream-server-go/authServer/internal/service"
	"log"
)

func main() {
	dbConn, err := db.OpenConn()
	if err != nil {
		log.Fatal("error connecting to database, err: ", err.Error())
	}
	defer log.Default().Println(dbConn.Close())

	// Init App
	keysRepo := repository.NewKeysRepository(dbConn)
	keysSvc := service.NewKeysService(keysRepo)
	keysHandler := handler.NewKeysHandler(keysSvc)

	server := echo.New()

	log.Default().Println("Routing application...")
	server.GET("/healthcheck", keysHandler.HealthCheck)
	server.POST("/auth", keysHandler.AuthStreaming)

	// Starting server
	server.Logger.Fatal(
		"Starting server at :8000",
		server.Start(
			":8000",
		),
	)
}
