package dto

import "github.com/google/uuid"

type TopUpReq struct {
	Amount float64   `json:"amount"`
	UserID uuid.UUID `json:"-"`
}
