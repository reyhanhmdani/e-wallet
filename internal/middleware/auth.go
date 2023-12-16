package middleware

import (
	"e-wallet/domain"
	"e-wallet/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strings"
)

func Authenticate(userService domain.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if token == "" {
			return ctx.SendStatus(401)
		}

		user, err := userService.ValidateTokenWithCache(ctx.Context(), token)
		if err != nil {
			logrus.Info("Received token:", token)
			logrus.Error("Error validating token: ", err)
			return ctx.SendStatus(util.GetHTTPStatus(err))
		}

		ctx.Locals("x-user", user)
		return ctx.Next()
	}
}

func AuthenticateJWT(userService domain.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")
		if tokenString == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Menghapus "Bearer " dari token string
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		userData, err := userService.ValidateTokenJwt(ctx.Context(), tokenString)
		if err != nil {
			logrus.Error(err)
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		// Set user data in locals
		ctx.Locals("x-user", userData)

		return ctx.Next()
	}
}
