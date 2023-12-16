package api

import (
	"context"
	"e-wallet/domain"
	"e-wallet/dto"
	"e-wallet/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

// untuk mengambil seluruh notification

type notificationAPI struct {
	notificationService domain.NotificationService
}

func NewNotification(app *fiber.App, authMid fiber.Handler, notificationService domain.NotificationService) {
	h := notificationAPI{
		notificationService: notificationService,
	}

	app.Get("/notification", authMid, h.GetUsersNotifications)
}

func (n notificationAPI) GetUsersNotifications(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 15*time.Second)
	defer cancel()

	user := ctx.Locals("x-user").(dto.UserData)
	logrus.Info(user)

	notification, err := n.notificationService.FindByUser(c, user.ID)
	if err != nil {
		logrus.Error(err)
		return ctx.Status(util.GetHTTPStatus(err)).JSON(dto.Response{
			Status:  util.GetHTTPStatus(err),
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(notification)

}
