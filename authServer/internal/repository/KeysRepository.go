package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/rmndvngrpslhr/stream-server-go/authServer/internal/model"
)

type KeysRepository interface {
	FindStreamKey(name, key string) (*model.Keys, error)
}

type keysRepository struct {
	DB *sql.DB
}

func NewKeysRepository(db *sql.DB) KeysRepository {
	return &keysRepository{DB: db}
}

var ErrQuery = errors.New("error finding stream key")

func (kr *keysRepository) FindStreamKey(name, key string) (*model.Keys, error) {
	fmt.Println(": ======================: looking for: \n", name, key)
	keys := &model.Keys{}
	row := kr.DB.QueryRow(
		`SELECT * FROM "lives" WHERE "name"=$1 AND "stream_key"=$2`,
		name, key,
	)

	if err := row.Scan(&keys.Name, &keys.KeyUUID); err != nil {
		log.Error(err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return &model.Keys{}, nil
		}
		return nil, ErrQuery
	}

	return keys, nil
}
