package service

import (
	"e-wallet/domain"
	"e-wallet/internal/config"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// lebih bagus pakai coreAPI => lebih customize dalam hal interface (dalam hal tampilan)
// kalau snap ini sudah di sediakan tempaltenya oleh midtrans itu sendiri

type midtransService struct {
	//client snap.Client
	config config.Midtrans
	envi   midtrans.EnvironmentType
}

func NewMidtrans(cnf *config.Config) domain.MidtransService {
	//var client snap.Client
	envi := midtrans.Sandbox
	if cnf.Midtrans.IsProd {
		envi = midtrans.Production
	}
	//client.New(cnf.Midtrans.Key, envi)

	return &midtransService{
		//client: client,
		config: cnf.Midtrans,
		envi:   envi,
	}
}

func (m midtransService) GenerateSnapURL(c context.Context, topUp *domain.TopUp) error {
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  topUp.ID,
			GrossAmt: int64(topUp.Amount),
		},
	}

	var client snap.Client
	client.New(m.config.Key, m.envi)
	snapResp, err := client.CreateTransaction(req)
	if err != nil {
		logrus.Error("error di bagain generate Snap url service", err)
		return err
	}
	topUp.SnapURL = snapResp.RedirectURL
	return nil
}

func (m midtransService) VerifyPayment(c context.Context, orderId string) (bool, error) {
	var client coreapi.Client
	//envi := midtrans.Sandbox
	//if m.config.IsProd {
	//	envi = midtrans.Production
	//}

	client.New(m.config.Key, m.envi)

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := client.CheckTransaction(orderId)
	if e != nil {
		return false, e
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					// TODO set transaction status on your database to 'success'
					return true, nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				// TODO set transaction status on your databaase to 'success'
				return true, nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}
	return false, nil
}
