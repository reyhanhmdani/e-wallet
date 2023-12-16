package domain

import (
	"context"
	"e-wallet/dto"
	"github.com/google/uuid"
)

type Factor struct {
	ID     uuid.UUID `db:"id"`
	UserId uuid.UUID `db:"user_id"`
	Pin    string    `db:"pin"`
}

type FactorRepository interface {
	FindByUser(ctx context.Context, id uuid.UUID) (Factor, error)
}

type FactorService interface {
	ValidatePin(ctx context.Context, req dto.ValidatePinReq) error
}
