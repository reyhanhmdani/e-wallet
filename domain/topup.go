package domain

import (
	"context"
	"e-wallet/dto"
	"github.com/google/uuid"
)

type TopUp struct {
	ID      string    `db:"id"`
	UserId  uuid.UUID `db:"user_id"`
	Status  int8      `db:"status"`
	Amount  float64   `db:"amount"`
	SnapURL string    `db:"snap_url"`
}

type TopUpRepository interface {
	FindById(c context.Context, id string) (TopUp, error)
	Insert(c context.Context, t *TopUp) error
	Update(c context.Context, t *TopUp) error
}

type TopUpService interface {
	InitializeTopUp(ctx context.Context, req dto.TopUpReq) (dto.TopUpRes, error)
	ConfirmedTopUp(ctx context.Context, id string) error
}
