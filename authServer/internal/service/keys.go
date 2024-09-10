package service

import (
	"github.com/rmndvngrpslhr/stream-server-go/authServer/internal/model"
	"github.com/rmndvngrpslhr/stream-server-go/authServer/internal/repository"
)

type KeyService interface {
	AuthStreamingKey(name, key string) (*model.Keys, error)
}

type keyService struct {
	keysRepository repository.KeysRepository
}

func NewKeysService(repo repository.KeysRepository) KeyService {
	return &keyService{keysRepository: repo}
}

func (ks *keyService) AuthStreamingKey(name, key string) (*model.Keys, error) {
	return ks.keysRepository.FindStreamKey(name, key)
}
