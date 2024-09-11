package service

import (
	"github.com/labstack/echo"
	"log"
	"path/filepath"
)

func ServeStream() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		streamName := ctx.Param("live")
		filePath := ctx.Param("*")

		if filePath == "" {
			filePath = "index.m3u8"
		}
		fileStreamPath := filepath.Join("/hls/live/", streamName, filePath)
		log.Default().Println("Stream file requested at: " + fileStreamPath)

		return ctx.File(fileStreamPath)
	}
}
