package jwttoken

import (
	"strconv"
	"time"
	app_err "web_test01/utility/errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecret = "web_test01_secret"
)

func GenerateJwtToken(subject int64) (string, error) {
	// create jwt token claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(int(subject)),
		"iss": "web_test01",
		"aud": "user",
		"exp": time.Now().Add(10 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
	})

	// create the signed jwt token
	token, err := claims.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func VerifyJwtToken(token string) (*jwt.MapClaims, error) {
	// instantiate claims object
	var claims jwt.MapClaims

	// get the parsed token along with claims
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// check if the token is valid
	if !parsedToken.Valid {
		return nil, app_err.ErrInvalidJwtToken
	}

	return &claims, nil
}
