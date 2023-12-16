package api

import (
	"e-wallet/domain"
	"e-wallet/dto"
	"e-wallet/internal/util"
	"github.com/gofiber/fiber/v2"
)

type topUpApi struct {
	topUpService domain.TopUpService
}

func NewTopUp(app *fiber.App, authmid fiber.Handler, topUpService domain.TopUpService) {
	t := topUpApi{
		topUpService: topUpService,
	}

	app.Post("topup/initialize", authmid, t.InitializeTopUp)
}

func (t topUpApi) InitializeTopUp(ctx *fiber.Ctx) error {
	var req dto.TopUpReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	user := ctx.Locals("x-user").(dto.UserData)
	req.UserID = user.ID

	res, err := t.topUpService.InitializeTopUp(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetHTTPStatus(err)).JSON(dto.Response{
			Status:  util.GetHTTPStatus(err),
			Message: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.SuccessTNData{
		Message: "Success top up",
		Status:  fiber.StatusOK,
		Data:    res,
	})
}
