package util

import (
	"e-wallet/domain"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func GetHTTPStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed),
		errors.Is(err, domain.ErrEmailNotVerified):
		return fiber.StatusUnauthorized
	case errors.Is(err, domain.UsernameOrEmailTaken),
		errors.Is(err, domain.OtpInvalid),
		errors.Is(err, domain.EmailAlreadyVerified),
		errors.Is(err, domain.InsufficientBalance),
		errors.Is(err, domain.SelfTransfer),
		errors.Is(err, domain.PinInvalid):
		return fiber.StatusBadRequest
	case errors.Is(err, domain.InquiryNotFound),
		errors.Is(err, domain.AccountNotFound),
		errors.Is(err, domain.UserNotFound),
		errors.Is(err, domain.TemplNotFound):
		return fiber.StatusNotFound
	default:
		return fiber.StatusInternalServerError
	}
}
