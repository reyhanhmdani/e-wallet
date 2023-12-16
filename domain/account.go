package domain

import (
	"context"
	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID `db:"id"`
	UserId        uuid.UUID `db:"user_id"`
	AccountNumber string    `db:"account_number"`
	Balance       float64   `db:"balance"`
}

type AccountRepository interface {
	FindUserById(ctx context.Context, id uuid.UUID) (Account, error)
	FindByAccountNumber(ctx context.Context, accNumber string) (Account, error)
	Update(ctx context.Context, account *Account) error
}
