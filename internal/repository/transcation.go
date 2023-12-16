package repository

import (
	"database/sql"
	"e-wallet/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type transactionRepository struct {
	db *goqu.Database
}

func NewTransaction(con *sql.DB) domain.TransactionRepository {
	return &transactionRepository{
		db: goqu.New("default", con),
	}
}

func (t transactionRepository) Insert(c context.Context, Transaction *domain.Transaction) error {
	exec := t.db.Insert("transactions").Rows(goqu.Record{
		"account_id":            Transaction.AccountId,
		"sof_number":            Transaction.SofNumber,
		"dof_number":            Transaction.DofNumber,
		"transaction_type":      Transaction.TransactionType,
		"amount":                Transaction.Amount,
		"transactions_datetime": Transaction.TransactionDatetime,
	}).Returning("id").Executor()
	_, err := exec.ScanStructContext(c, Transaction)
	if err != nil {
		logrus.Error("Insert transaction failed: ", err)
		return err
	}

	return err
}
