package api

import (
	"e-wallet/domain"
	"e-wallet/dto"
	"github.com/gofiber/fiber/v2"
)

type midtransApi struct {
	midtransService domain.MidtransService
	topUpService    domain.TopUpService
}

func NewMidtrans(app *fiber.App, midtransServce domain.MidtransService, topUpService domain.TopUpService) {
	m := midtransApi{
		midtransService: midtransServce,
		topUpService:    topUpService,
	}

	app.Post("midtrans/payment-callback", m.paymentHanlerNotification)
}

func (m midtransApi) paymentHanlerNotification(ctx *fiber.Ctx) error {
	var notificationPayload map[string]interface{}
	if err := ctx.BodyParser(&notificationPayload); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	// 3. Get order-id from payload
	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	succcess, _ := m.midtransService.VerifyPayment(ctx.Context(), orderId)
	if succcess {
		_ = m.topUpService.ConfirmedTopUp(ctx.Context(), orderId)
	}
	return ctx.Status(fiber.StatusBadRequest).JSON(dto.SuccessCommonRes{
		Message: "notification handler gagal",
		Status:  fiber.StatusBadRequest,
	})
}
