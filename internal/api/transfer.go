package api

import (
	"e-wallet/domain"
	"e-wallet/dto"
	"e-wallet/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type transferApi struct {
	transactionService domain.TransactionService
	factorService      domain.FactorService
}

func NewTransfer(app *fiber.App,
	authMid fiber.Handler,
	transactionService domain.TransactionService,
	factorService domain.FactorService) {
	h := &transferApi{
		transactionService: transactionService,
		factorService:      factorService,
	}

	app.Post("transfer/inquiry", authMid, h.TransferInquiry)
	app.Post("transfer/execute", authMid, h.TransferExecute)
}

func (t transferApi) TransferInquiry(ctx *fiber.Ctx) error {
	var req dto.TransferInquiryReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	inquiry, err := t.transactionService.TransferInquiry(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetHTTPStatus(err)).JSON(dto.Response{
			Status:  util.GetHTTPStatus(err),
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.SuccessTNData{
		Message: "Success Generate Code",
		Status:  fiber.StatusOK,
		Data:    inquiry,
	})
}
func (t transferApi) TransferExecute(ctx *fiber.Ctx) error {
	var req dto.TransferExecuteReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user := ctx.Locals("x-user").(dto.UserData)
	if err := t.factorService.ValidatePin(ctx.Context(), dto.ValidatePinReq{
		PIN:    req.PIN,
		UserID: user.ID,
	}); err != nil {
		logrus.Error("kesalahan validate pin ", err)
		return ctx.Status(util.GetHTTPStatus(err)).JSON(dto.Response{
			Status:  util.GetHTTPStatus(err),
			Message: "salah pin nya",
		})
	}

	err := t.transactionService.TransferExecute(ctx.Context(), req)
	if err != nil {
		logrus.Error(err)
		return ctx.Status(util.GetHTTPStatus(err)).JSON(dto.Response{
			Status:  util.GetHTTPStatus(err),
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.SuccessCommonRes{
		Message: "Success mengirim ke no tujuan",
		Status:  fiber.StatusOK,
	})
}
