package domain

import (
	"context"
	"e-wallet/dto"
)

type IpCheckerService interface {
	Query(ctx context.Context, ip string) (dto.IpChecker, error)
}
