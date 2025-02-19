package auth

import (
	"strings"
	app_err "web_test01/utility/errors"
	"web_test01/utility/response"
	jwttoken "web_test01/utility/token"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	reqHeaders := c.GetReqHeaders()

	tokenVal, ok := reqHeaders["Authorization"]
	if !ok {
		return response.ErrorJson(c, fiber.StatusUnauthorized, app_err.ErrAuthTokenMissing, nil)
	}

	if !strings.Contains(tokenVal[0], "Bearer") {
		return response.ErrorJson(c, fiber.StatusUnauthorized, app_err.ErrAuthTokenMissing, nil)
	}

	auth := strings.Split(tokenVal[0], " ")

	claims, err := jwttoken.VerifyJwtToken(auth[1])
	if err != nil {
		return response.ErrorJson(c, fiber.StatusUnauthorized, app_err.ErrAuthTokenInvalid, nil)
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return response.ErrorJson(c, fiber.StatusUnauthorized, app_err.ErrAuthTokenInvalid, nil)
	}

	c.Locals("UserId", sub)

	return c.Next()
}
