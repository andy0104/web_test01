package services

import (
	"web_test01/storage"
	"web_test01/types"

	"go.uber.org/zap"
)

type Services struct {
	User interface {
		RegisterUser(types.RegisterPayload) (string, error)
		LoginUser(types.LoginPayload) (string, error)
		GetUserInfoById(int64) (*storage.User, error)
	}
}

func NewServices(
	logger *zap.SugaredLogger,
	store storage.Storage,
) Services {
	return Services{
		User: &User{logger: logger, store: store},
	}
}
