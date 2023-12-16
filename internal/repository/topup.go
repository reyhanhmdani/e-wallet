package repository

import (
	"context"
	"database/sql"
	"e-wallet/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/sirupsen/logrus"
)

type topUpRepository struct {
	db *goqu.Database
}

func NewTopUp(conn *sql.DB) domain.TopUpRepository {
	return &topUpRepository{
		db: goqu.New("default", conn),
	}
}

func (t topUpRepository) FindById(c context.Context, id string) (topUp domain.TopUp, err error) {
	dataset := t.db.From("topup").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(c, &topUp)
	if err != nil {
		logrus.Error("error di bagian repo find By Id", err)
		return domain.TopUp{}, err
	}
	return
}

func (t topUpRepository) Insert(c context.Context, topUp *domain.TopUp) error {
	executor := t.db.Insert("topup").Rows(goqu.Record{
		"id":       topUp.ID,
		"user_id":  topUp.UserId,
		"amount":   topUp.Amount,
		"status":   topUp.Status,
		"snap_url": topUp.SnapURL,
	}).Executor()
	_, err := executor.ExecContext(c)
	if err != nil {
		logrus.Error("error di bagian repo insert", err)
	}
	return err
}

func (t topUpRepository) Update(c context.Context, topUp *domain.TopUp) error {
	executor := t.db.Update("topup").Where(goqu.Ex{
		"id": topUp.ID,
	}).Set(goqu.Record{
		"amount":   topUp.Amount,
		"status":   topUp.Status,
		"snap_url": topUp.SnapURL,
	}).Executor()
	_, err := executor.ExecContext(c)
	if err != nil {
		logrus.Error("error di bagian repo insert", err)
	}
	return err
}
