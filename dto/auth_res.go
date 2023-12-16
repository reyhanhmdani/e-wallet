package dto

import "github.com/google/uuid"

type AuthRes struct {
	Token  string    `json:"token"`
	UserID uuid.UUID `json:"-"`
}
