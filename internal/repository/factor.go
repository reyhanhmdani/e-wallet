package repository

import (
	"context"
	"database/sql"
	"e-wallet/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type factorRepository struct {
	db *goqu.Database
}

func NewFactor(conn *sql.DB) domain.FactorRepository {
	return &factorRepository{
		db: goqu.New("default", conn),
	}
}

func (f factorRepository) FindByUser(ctx context.Context, id uuid.UUID) (factor domain.Factor, err error) {
	dataset := f.db.From("factors").Where(goqu.Ex{
		"user_id": id,
	})
	_, err = dataset.ScanStructContext(ctx, &factor)
	if err != nil {
		logrus.Error("Id nya ga dapat (di repo factor)", err)
	}
	return
}
