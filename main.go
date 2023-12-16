package main

import (
	"e-wallet/dto"
	"e-wallet/internal/api"
	"e-wallet/internal/component"
	"e-wallet/internal/config"
	"e-wallet/internal/middleware"
	"e-wallet/internal/repository"
	"e-wallet/internal/service"
	"e-wallet/internal/sse"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func main() {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//logrus.SetLevel(logrus.InfoLevel)

	cnf := config.Get()
	dbConnection := component.GetDbConnection(cnf)
	//cacheConnection := component.GetCacheConnection()
	cacheConnection := repository.NewRedisClient(cnf)
	hub := &dto.Hub{NotificationChannel: map[uuid.UUID]chan dto.NotificationData{}}

	component.Log.Info("Applications Hello world")

	userRepository := repository.NewUser(dbConnection)
	accountRepository := repository.NewAccount(dbConnection)
	transactionRepository := repository.NewTransaction(dbConnection)
	notificationRepository := repository.NewNotification(dbConnection)
	templateRepository := repository.NewTemplate(dbConnection)
	topUpRepository := repository.NewTopUp(dbConnection)
	factorRepository := repository.NewFactor(dbConnection)
	loginLogRepository := repository.NewLoginLog(dbConnection)

	queueService := service.NewQueue(cnf)
	emailService := service.NewEmail(cnf, queueService)
	userService := service.NewUser(userRepository, cacheConnection, emailService)
	notificationService := service.NewNotification(notificationRepository, templateRepository, hub)
	transactionService := service.NewTransaction(accountRepository, transactionRepository, cacheConnection, notificationService)
	midtransService := service.NewMidtrans(cnf)
	topUpService := service.NewTopUp(notificationService, midtransService, topUpRepository, accountRepository, transactionRepository)
	factorService := service.NewFactor(factorRepository)
	ipCheckerService := service.NewIpChecker()
	fdsService := service.NewFds(ipCheckerService, loginLogRepository)

	authMiddleware := middleware.AuthenticateJWT(userService)

	app := fiber.New()
	api.NewAuth(app, userService, authMiddleware, fdsService)
	api.NewTransfer(app, authMiddleware, transactionService, factorService)
	api.NewNotification(app, authMiddleware, notificationService)
	api.NewTopUp(app, authMiddleware, topUpService)
	api.NewMidtrans(app, midtransService, topUpService)

	sse.NewNotification(app, authMiddleware, hub)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}

// q1PqLc0r4vLBUBmX8WKu
