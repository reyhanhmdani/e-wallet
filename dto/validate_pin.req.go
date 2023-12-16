package dto

import "github.com/google/uuid"

type ValidatePinReq struct {
	PIN    string
	UserID uuid.UUID
}
