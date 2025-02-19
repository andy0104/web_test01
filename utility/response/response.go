package response

import "github.com/gofiber/fiber/v2"

type jsonResponse struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

type errorResponse struct {
	Message string `json:"message"`
	Details any    `json:"details"`
}

func WriteJson(c *fiber.Ctx, status int, data any, isError bool) error {
	resp := new(jsonResponse)
	if isError {
		resp.Data = nil
		resp.Error = data
	} else {
		resp.Data = data
		resp.Error = nil
	}
	return c.Status(status).JSON(resp)
}

func ErrorJson(c *fiber.Ctx, status int, err error, data any) error {
	errors := errorResponse{
		Message: err.Error(),
		Details: data,
	}
	return WriteJson(c, status, errors, true)
}
