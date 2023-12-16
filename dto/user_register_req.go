package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type UserRegisterReq struct {
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateReq struct {
	UserRegisterReq
	Fullname string `json:"fullname" validate:"required" create:"true"`
	Password string `json:"password" validate:"required,min=1" create:"true"`
	Username string `json:"username" validate:"required" create:"true"`
}

func ValidateUserRegisterReq(req UserCreateReq) error {
	if err := validate.Struct(req); err != nil {
	}
	return fiber.NewError(fiber.StatusBadRequest, "error mamang")
}
