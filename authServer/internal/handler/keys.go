package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/rmndvngrpslhr/stream-server-go/authServer/internal/model"
	"github.com/rmndvngrpslhr/stream-server-go/authServer/internal/service"
	"io"
	"log"
	"net/http"
	"strings"
)

type KeysHandler interface {
	AuthStreaming(context echo.Context) error
	HealthCheck(context echo.Context) error
}

type keysHandler struct {
	keysService service.KeyService
}

func NewKeysHandler(service service.KeyService) KeysHandler {
	return &keysHandler{keysService: service}
}

func (kh *keysHandler) AuthStreaming(context echo.Context) error {
	log.Default().Println("Running authentication...")
	body := context.Request().Body

	fields, _ := io.ReadAll(body)
	log.Default().Println("Auth...\n", string(fields))

	givenKeyValues := getStreamKeys(fields)
	keys, err := kh.keysService.AuthStreamingKey(givenKeyValues.Name, givenKeyValues.KeyUUID)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}

	if keys.KeyUUID == "" {
		return context.String(http.StatusForbidden, "user not authenticated")
	}

	log.Default().Println("user authenticated")
	newStreamURL := fmt.Sprintf("rtmp:127.0.0.1:1935/hls-live/%s", keys.Name)
	log.Default().Println("Redirecting to: ", newStreamURL)
	return context.Redirect(http.StatusFound, newStreamURL)
}

func getStreamKeys(s []byte) model.Keys {
	var authValues model.Keys

	pairs := strings.Split(string(s), "&")
	for _, pair := range pairs {
		splitPair := strings.Split(pair, "=")
		key := splitPair[0]
		value := splitPair[1]

		if key == "name" {
			allValues := strings.Split(value, "_")

			authValues.Name = allValues[0]
			authValues.KeyUUID = allValues[1]
		}
	}
	return authValues
}

func (kh *keysHandler) HealthCheck(context echo.Context) error {
	return context.String(http.StatusOK, "Auth App - STATUS: HEALTHY")
}
