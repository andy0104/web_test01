package storage

import (
	"web_test01/types"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Storage struct {
	User interface {
		Create(types.RegisterPayload) (int64, error)
		GetUserByEmail(string) (*User, error)
		GetUserById(int64) (*User, error)
	}
}

func NewStorage(db *sqlx.DB, logger *zap.SugaredLogger) Storage {
	return Storage{
		User: &UserStore{db: db, logger: logger},
	}
}
