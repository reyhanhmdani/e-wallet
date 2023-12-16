package service

import (
	"context"
	"e-wallet/domain"
	"e-wallet/dto"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type TopUpService struct {
	notificationService   domain.NotificationService // buat mengirim notifikasi ke dalam yang nerima / berhasil topup
	midtransService       domain.MidtransService
	topUpRepository       domain.TopUpRepository
	accountRepository     domain.AccountRepository
	transactionRepository domain.TransactionRepository
}

func NewTopUp(notificationService domain.NotificationService,
	midtransService domain.MidtransService,
	topUpRepository domain.TopUpRepository,
	accountRepository domain.AccountRepository,
	transactionRepository domain.TransactionRepository) domain.TopUpService {
	return TopUpService{
		notificationService:   notificationService,
		topUpRepository:       topUpRepository,
		midtransService:       midtransService,
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

func (t TopUpService) InitializeTopUp(ctx context.Context, req dto.TopUpReq) (dto.TopUpRes, error) {
	// kita minta ke midtrans buat generate
	topUp := domain.TopUp{
		ID:     uuid.NewString(),
		UserId: req.UserID,
		Status: 0,
		Amount: req.Amount,
		//SnapURL: "",
	}
	// generate snap url
	err := t.midtransService.GenerateSnapURL(ctx, &topUp)
	if err != nil {
		return dto.TopUpRes{}, err
	}

	err = t.topUpRepository.Insert(ctx, &topUp)
	if err != nil {
		return dto.TopUpRes{}, err
	}

	return dto.TopUpRes{
		SnapUrl: topUp.SnapURL,
	}, nil
}

func (t TopUpService) ConfirmedTopUp(ctx context.Context, id string) error {
	// kita melakukan updated
	// cari dlu top up nya, bener ada ga tu kira kira top up nya
	topUp, err := t.topUpRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if topUp == (domain.TopUp{}) {
		return domain.TopUpReqNotFound
	}

	account, err := t.accountRepository.FindUserById(ctx, topUp.UserId)
	if err != nil {
		logrus.Error("error di bagian service top up ", err)
		return err
	}
	if account == (domain.Account{}) {
		return domain.AccountNotFound
	}

	// topup darimana nih
	err = t.transactionRepository.Insert(ctx, &domain.Transaction{
		AccountId:           account.ID,
		SofNumber:           "00",
		DofNumber:           account.AccountNumber,
		TransactionType:     "C",
		Amount:              topUp.Amount,
		TransactionDatetime: time.Now(),
	})
	if err != nil {
		logrus.Error("error di bagian Insert topup service top up ", err)
		return err
	}

	account.Balance += topUp.Amount
	err = t.accountRepository.Update(ctx, &account)
	if err != nil {
		logrus.Error("error di bagian update service top up ", err)
		return err
	}

	data := map[string]string{
		"amount": fmt.Sprintf("%2.f", topUp.Amount),
	}
	err = t.notificationService.Insert(ctx, account.UserId, "TOPUP_SUCCESS", data)
	if err != nil {
		logrus.Error("error di bagian Insert notif, service top up ", err)
		return err
	}

	return nil
}
