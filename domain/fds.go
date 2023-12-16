package domain

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

// Fraud Detection System (FDS)

type FdsService interface {
	// check apakah login nya berpotensi fraud apa tidak
	IsAuthorized(ctx context.Context, ip string, userId uuid.UUID) bool
}
