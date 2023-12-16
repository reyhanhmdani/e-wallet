package api

import (
	"e-wallet/domain"
	"e-wallet/dto"
	"e-wallet/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type authAPI struct {
	userService domain.UserService
	fdsService  domain.FdsService
}

func NewAuth(app *fiber.App, userService domain.UserService, authMiddleware fiber.Handler, fdsService domain.FdsService) {
	h := authAPI{
		userService: userService,
		fdsService:  fdsService,
	}

	app.Post("token/login", h.GenerateToken)
	app.Get("token/validate", authMiddleware, h.ValidateToken)
	app.Get("user", h.GetUsers)
	app.Post("user/register", h.RegisterUser)
	app.Post("user/validateOtp", h.ValidateOtp)
}

func (a authAPI) GetUsers(ctx *fiber.Ctx) error {
	users, err := a.userService.AllUsers(ctx.Context())
	if err != nil {

	}
	return ctx.Status(fiber.StatusOK).JSON(dto.SuccessTNData{
		Message: "Success get all user",
		Status:  fiber.StatusOK,
		Data:    users,
	})

}

func (a authAPI) GenerateToken(ctx *fiber.Ctx) error {
	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	token, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {
		logrus.Error(err)
		return ctx.Status(util.GetHTTPStatus(err)).JSON(dto.Response{
			Status:  util.GetHTTPStatus(err),
			Message: err.Error(),
		})
	}
	if !a.fdsService.IsAuthorized(ctx.Context(), ctx.Get("X-FORWARDED-FOR"), token.UserID) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.Response{
			Status:  fiber.StatusUnauthorized,
			Message: "Is authorized gagal",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.SuccessTNData{
		Message: "Success Login",
		Status:  fiber.StatusOK,
		Data:    token,
	})
}

func (a authAPI) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")
	logrus.Info("Authenticated user in handler:", user)
	return ctx.Status(200).JSON(user)
}

func (a authAPI) RegisterUser(ctx *fiber.Ctx) error {
	var request dto.UserCreateReq
	if err := ctx.BodyParser(&request); err != nil {
		logrus.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	res, err := a.userService.Register(ctx.Context(), request)
	if err != nil {
		logrus.Error(err)
		return ctx.Status(util.GetHTTPStatus(err)).JSON(dto.Response{
			Status:  util.GetHTTPStatus(err),
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.SuccessTNData{
		Message: "Success Register dan tolong validasi Emailnya terlebih dahulu",
		Status:  fiber.StatusCreated,
		Data:    res,
	})
}

func (a authAPI) ValidateOtp(ctx *fiber.Ctx) error {
	var request dto.ValidateOTPReq
	if err := ctx.BodyParser(&request); err != nil {
		logrus.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.Response{
			Status:  fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err := a.userService.ValidateOTP(ctx.Context(), request)
	if err != nil {
		logrus.Error("otp nya salah : ", err)
		return ctx.Status(util.GetHTTPStatus(err)).JSON(dto.Response{
			Status:  util.GetHTTPStatus(err),
			Message: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(dto.SuccessCommonRes{
		Message: "Email anda sudah terverifikasi",
		Status:  fiber.StatusCreated,
	})
}
