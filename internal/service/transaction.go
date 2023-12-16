package service

import (
	"e-wallet/domain"
	"e-wallet/dto"
	"e-wallet/internal/component"
	"e-wallet/internal/util"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type transactionService struct {
	accountRepository     domain.AccountRepository
	transactionRepository domain.TransactionRepository
	cacheRepository       domain.CacheRepository
	notificationService   domain.NotificationService
}

func NewTransaction(accountRepository domain.AccountRepository,
	transactionrRepository domain.TransactionRepository,
	cacheRepository domain.CacheRepository,
	notificationService domain.NotificationService) domain.TransactionService {
	return &transactionService{
		accountRepository:     accountRepository,
		transactionRepository: transactionrRepository,
		cacheRepository:       cacheRepository,
		notificationService:   notificationService,
	}
}

func (t transactionService) TransferInquiry(c context.Context, req dto.TransferInquiryReq) (dto.TransferInquiryRes, error) {
	user := c.Value("x-user").(dto.UserData)
	myAccount, err := t.accountRepository.FindUserById(c, user.ID)
	if err != nil {
		return dto.TransferInquiryRes{}, err
	}
	if myAccount == (domain.Account{}) {
		return dto.TransferInquiryRes{}, domain.AccountNotFound
	}
	// check destination
	dofAccount, err := t.accountRepository.FindByAccountNumber(c, req.AccountNumber)
	if err != nil {
		return dto.TransferInquiryRes{}, err
	}
	if dofAccount == (domain.Account{}) {
		return dto.TransferInquiryRes{}, domain.AccountNotFound
	}

	// check self-transfer
	if myAccount.AccountNumber == dofAccount.AccountNumber {
		return dto.TransferInquiryRes{}, domain.SelfTransfer
	}

	if myAccount.Balance < req.Amount {
		return dto.TransferInquiryRes{}, domain.InsufficientBalance
	}

	inquiryKey := util.GenerateRandomString(32)

	jsonData, err := json.Marshal(req)
	_ = t.cacheRepository.Set(inquiryKey, jsonData)

	return dto.TransferInquiryRes{
		InquiryKey: inquiryKey,
	}, nil

}

func (t transactionService) TransferExecute(c context.Context, req dto.TransferExecuteReq) error {
	component.Log.Info("Starting execute transfer")
	component.Log.Debugf("%s to %s", req)

	// ambil cache nya
	val, err := t.cacheRepository.Get(req.InquiryKey)
	if err != nil {
		return domain.InquiryNotFound
	}

	var reqInq dto.TransferInquiryReq
	// decode val dari cachenya
	_ = json.Unmarshal(val, &reqInq)

	user := c.Value("x-user").(dto.UserData)
	myAccount, err := t.accountRepository.FindUserById(c, user.ID)
	if err != nil {
		return err
	}
	// check destination
	dofAccount, err := t.accountRepository.FindByAccountNumber(c, reqInq.AccountNumber)
	if err != nil {
		return err
	}

	// cek apakah akun sumber dan tujuan sama
	if myAccount.AccountNumber == dofAccount.AccountNumber {
		return domain.SelfTransfer
	}

	component.Log.Debugf("%s to %s", myAccount.AccountNumber, dofAccount.AccountNumber)
	debitTransaction := domain.Transaction{
		AccountId:           myAccount.ID,
		SofNumber:           myAccount.AccountNumber,
		DofNumber:           dofAccount.AccountNumber,
		TransactionType:     "D",
		Amount:              reqInq.Amount,
		TransactionDatetime: time.Now(),
	}

	err = t.transactionRepository.Insert(c, &debitTransaction)
	if err != nil {
		logrus.Error("insert debit Transaction gagal")
		return err
	}

	// memasukkan uang ke akun tujuannya
	creditTransaction := domain.Transaction{
		AccountId:           myAccount.ID,
		SofNumber:           myAccount.AccountNumber,
		DofNumber:           dofAccount.AccountNumber,
		TransactionType:     "C",
		Amount:              reqInq.Amount,
		TransactionDatetime: time.Now(),
	}

	err = t.transactionRepository.Insert(c, &creditTransaction)
	if err != nil {
		logrus.Error("insert credit Transaction gagal")
		return err
	}

	// pengurangan dominal
	myAccount.Balance -= reqInq.Amount
	err = t.accountRepository.Update(c, &myAccount)
	if err != nil {
		logrus.Error("update my account err", err)
		return err
	}

	dofAccount.Balance += reqInq.Amount
	err = t.accountRepository.Update(c, &dofAccount)
	if err != nil {
		logrus.Error("update dof acount err", err)
		return err
	}

	go t.notificationAfterTransfer(myAccount, dofAccount, reqInq.Amount)

	return nil

}

func (t transactionService) notificationAfterTransfer(sofAccount domain.Account, dofAccount domain.Account, amount float64) {
	data := map[string]string{
		"amount": fmt.Sprintf("%2.f", amount),
	}

	_ = t.notificationService.Insert(context.Background(), sofAccount.UserId, "TRANSFER", data)
	_ = t.notificationService.Insert(context.Background(), dofAccount.UserId, "TRANSFER_DES", data)
	//
	//notificationSender := domain.Notification{
	//	UserId:    sofAccount.UserId,
	//	Title:     "Transfer berhasil",
	//	Body:      fmt.Sprintf("Transfer Senilai %.2f berhasil", amount),
	//	IsRead:    0,
	//	Status:    1,
	//	CreatedAt: time.Now(),
	//}
	//notificationReceiver := domain.Notification{
	//	UserId:    dofAccount.UserId,
	//	Title:     "Dana di terima",
	//	Body:      fmt.Sprintf("Dana di terima senilai %.2f", amount),
	//	Status:    1,
	//	IsRead:    0,
	//	CreatedAt: time.Now(),
	//}
	//_ = t.notificationRepository.Insert(context.Background(), &notificationSender)
	//// mengirim data notificationnya
	//if channel, ok := t.hub.NotificationChannel[sofAccount.UserId]; ok {
	//	channel <- dto.NotificationData{
	//		Id:        notificationSender.Id,
	//		Title:     notificationSender.Title,
	//		Body:      notificationSender.Body,
	//		Status:    notificationSender.Status,
	//		IsRead:    notificationSender.IsRead,
	//		CreatedAt: notificationSender.CreatedAt,
	//	}
	//}
	//
	//_ = t.notificationRepository.Insert(context.Background(), &notificationReceiver)
	//// mengirim data notificationnya
	//if channel, ok := t.hub.NotificationChannel[dofAccount.UserId]; ok {
	//	channel <- dto.NotificationData{
	//		Id:        notificationReceiver.Id,
	//		Title:     notificationReceiver.Title,
	//		Body:      notificationReceiver.Body,
	//		Status:    notificationReceiver.Status,
	//		IsRead:    notificationReceiver.IsRead,
	//		CreatedAt: notificationReceiver.CreatedAt,
	//	}
	//}
}
