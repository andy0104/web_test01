package app_err

import "errors"

var (
	ErrNoRecords        = errors.New("resource not found")
	ErrInvalidLogin     = errors.New("invalid login credentials")
	ErrInternalServer   = errors.New("internal server error")
	ErrEmailInUseServer = errors.New("email is already in use")
	ErrInvalidJwtToken  = errors.New("invalid token")
	ErrAuthTokenMissing = errors.New("missing authorization token")
	ErrAuthTokenInvalid = errors.New("invalid authorization token")
)
