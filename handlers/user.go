package handlers

import (
	"errors"
	"fmt"
	"strconv"
	"web_test01/services"
	"web_test01/types"
	app_err "web_test01/utility/errors"
	"web_test01/utility/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UserHandler struct {
	logger   *zap.SugaredLogger
	services services.Services
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	body := new(types.RegisterPayload)

	if err := c.BodyParser(body); err != nil {
		uh.logger.Errorw("Register error", err)
		return response.ErrorJson(c, fiber.StatusBadRequest, err, nil)
	}

	if err := Validate.Struct(body); err != nil {
		uh.logger.Errorw("Register validation error", "error", err)
		// get detailed error messages
		errs := err.(validator.ValidationErrors)
		for e, k := range errs {
			fmt.Printf("%v %v\n", e, k)
		}
		return response.ErrorJson(c, fiber.StatusBadRequest, err, nil)
	}

	uh.logger.Infow("Body params", "body", body)

	resp, err := uh.services.User.RegisterUser(*body)
	if err != nil {
		uh.logger.Errorw("Register service error", err)
		switch {
		case errors.Is(err, app_err.ErrEmailInUseServer):
			return response.ErrorJson(c, fiber.StatusConflict, err, nil)
		default:
			return response.ErrorJson(c, fiber.StatusInternalServerError, err, nil)
		}
	}

	uh.logger.Infow("Register user resp", "resp", resp)
	return response.WriteJson(c, fiber.StatusCreated, fiber.Map{
		"message": resp,
	}, false)
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	body := new(types.LoginPayload)

	if err := c.BodyParser(body); err != nil {
		uh.logger.Errorw("Login error", err)
		return response.ErrorJson(c, fiber.StatusBadRequest, err, nil)
	}

	if err := Validate.Struct(body); err != nil {
		uh.logger.Errorw("Login validation error", err)
		return response.ErrorJson(c, fiber.StatusBadRequest, err, nil)
	}

	uh.logger.Infow("Login params", "body", body)

	resp, err := uh.services.User.LoginUser(*body)
	if err != nil {
		switch {
		case errors.Is(err, app_err.ErrNoRecords) || errors.Is(err, app_err.ErrInvalidLogin):
			return response.ErrorJson(c, fiber.StatusUnauthorized, err, nil)
		default:
			return response.ErrorJson(c, fiber.StatusInternalServerError, err, nil)
		}
	}

	return response.WriteJson(c, fiber.StatusOK, fiber.Map{
		"token": resp,
	}, false)
}

func (uh *UserHandler) Profile(c *fiber.Ctx) error {
	// get the userId from the fiber context
	localData := c.Locals("UserId")
	userId, _ := strconv.Atoi(localData.(string))
	user, err := uh.services.User.GetUserInfoById(int64(userId))
	if err != nil {
		switch {
		case errors.Is(err, app_err.ErrNoRecords):
			return response.ErrorJson(c, fiber.StatusNoContent, err, nil)
		default:
			return response.ErrorJson(c, fiber.StatusInternalServerError, err, nil)
		}
	}

	return response.WriteJson(c, fiber.StatusOK, fiber.Map{
		"user": user,
	}, false)
}

func (uh *UserHandler) Coverpic(c *fiber.Ctx) error {
	localData := c.Locals("UserId")
	userId, _ := strconv.Atoi(localData.(string))
	form, err := c.MultipartForm()
	if err != nil {
		return response.ErrorJson(c, fiber.StatusInternalServerError, err, nil)
	}

	file := form.File["profilepic"][0]
	fmt.Println(file.Filename, file.Size/1024, file.Header.Get("Content-Type"))
	ff, _ := file.Open()
	defer ff.Close()

	fmt.Println(ff)

	// get the file data
	// fileData, _ := io.ReadAll(ff)
	// fmt.Println(fileData)

	return response.WriteJson(c, fiber.StatusOK, fiber.Map{
		"user": int64(userId),
	}, false)
}
