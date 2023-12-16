package repository

import (
	"context"
	"database/sql"
	"e-wallet/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type accountRepository struct {
	db *goqu.Database
}

func NewAccount(con *sql.DB) domain.AccountRepository {
	return &accountRepository{
		db: goqu.New("default", con),
	}
}

func (a accountRepository) FindUserById(c context.Context, id uuid.UUID) (account domain.Account, err error) {
	dataset := a.db.From("accounts").Where(goqu.Ex{
		"user_id": id,
	})
	_, err = dataset.ScanStructContext(c, &account)
	return
}

func (a accountRepository) FindByAccountNumber(c context.Context, accNumber string) (account domain.Account, err error) {
	dataset := a.db.From("accounts").Where(goqu.Ex{
		"account_number": accNumber,
	})
	_, err = dataset.ScanStructContext(c, &account)
	return
}

func (a accountRepository) Update(c context.Context, account *domain.Account) error {
	executor := a.db.Update("accounts").Where(goqu.Ex{
		"id": account.ID,
	}).Set(goqu.Record{
		"balance": account.Balance,
	}).Executor()
	_, err := executor.ExecContext(c)
	return err
}
