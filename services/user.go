package services

import (
	"errors"
	"web_test01/storage"
	"web_test01/types"
	app_err "web_test01/utility/errors"
	"web_test01/utility/hasher"
	jwttoken "web_test01/utility/token"

	"go.uber.org/zap"
)

type User struct {
	logger *zap.SugaredLogger
	store  storage.Storage
}

func (u *User) RegisterUser(up types.RegisterPayload) (string, error) {
	// hash the password before storing
	hashed, err := hasher.HashText(up.Password)
	if err != nil {
		u.logger.Errorw("Password hash error", "err", err)
		return "", err
	}
	u.logger.Infow("Hash password", "hash", string(hashed), "hash len", len(string(hashed)))
	up.Password = string(hashed)

	// check if user email is already in use
	user, err := u.store.User.GetUserByEmail(up.Email)
	if err != nil {
		switch {
		case errors.Is(err, app_err.ErrNoRecords):
			break
		default:
			u.logger.Errorw("Get user by email err", "err", err)
			return "", err
		}
	}

	if user != nil {
		u.logger.Errorw("User email already exist", "err", app_err.ErrEmailInUseServer)
		return "", app_err.ErrEmailInUseServer
	}

	// create user
	resp, err := u.store.User.Create(up)
	if err != nil {
		return "", err
	}

	if resp == int64(0) {
		return "", app_err.ErrInternalServer
	}
	return "User created", nil
}

func (u *User) LoginUser(up types.LoginPayload) (string, error) {
	user, err := u.store.User.GetUserByEmail(up.Email)
	if err != nil {
		return "", err
	}

	if err := hasher.CompareHashText(user.Password, up.Password); err != nil {
		u.logger.Errorw("Compare password error", "err", err)
		return "", app_err.ErrInvalidLogin
	}

	// generate the jwt token
	token, err := jwttoken.GenerateJwtToken(user.ID)
	if err != nil {
		u.logger.Errorw("JWT token error", "err", err)
		return "", app_err.ErrInternalServer
	}

	return token, nil
}

func (u *User) GetUserInfoById(userId int64) (*storage.User, error) {
	user, err := u.store.User.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
