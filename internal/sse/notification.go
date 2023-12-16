package sse

import (
	"bufio"
	"e-wallet/dto"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type notificationSSE struct {
	hub *dto.Hub // memakai pointer supaya kalau ada perubahan apapun yang ada disini maka ber efek ke yang lain juga ..
}

func NewNotification(app *fiber.App, authmid fiber.Handler, hub *dto.Hub) {
	h := notificationSSE{
		hub: hub,
	}

	app.Get("sse/notification-stream", authmid, h.StreamNotification)
}

func (n notificationSSE) StreamNotification(ctx *fiber.Ctx) error { // jadi nanti user nya akan stream ke endpoint ini, untuk mendapatkan notif
	ctx.Set("Content-type", "text/event-stream")

	user := ctx.Locals("x-user").(dto.UserData)
	n.hub.NotificationChannel[user.ID] = make(chan dto.NotificationData)

	ctx.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		event := fmt.Sprintf("event: %s\n"+
			"data:\n\n", "initial")
		_, _ = fmt.Fprint(w, event)
		_ = w.Flush()

		for notification := range n.hub.NotificationChannel[user.ID] {
			data, _ := json.Marshal(notification)

			event = fmt.Sprintf("event: %s\n"+
				"data: %s\n\n", "notification-updated", data)

			_, _ = fmt.Fprint(w, event)
			_ = w.Flush()
		}
	})
	return nil
}
