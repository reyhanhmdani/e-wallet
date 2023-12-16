package domain

import (
	"context"
	"e-wallet/dto"
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID                  uuid.UUID `db:"id"`
	AccountId           uuid.UUID `db:"account_id"`
	SofNumber           string    `db:"sof_number"`
	DofNumber           string    `db:"dof_number"`
	TransactionType     string    `db:"transaction_type"`
	Amount              float64   `db:"amount"`
	TransactionDatetime time.Time `db:"transactions_datetime"`
}

type TransactionRepository interface {
	Insert(c context.Context, Transaction *Transaction) error
}

type TransactionService interface {
	//Transfer
	TransferInquiry(c context.Context, req dto.TransferInquiryReq) (dto.TransferInquiryRes, error)
	TransferExecute(c context.Context, req dto.TransferExecuteReq) error
}
