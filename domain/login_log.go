package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type LoginLog struct {
	Id           uuid.UUID `db:"id"`
	UserId       uuid.UUID `db:"user_id"`
	IsAuthorized bool      `db:"is_authorized"`
	IpAddress    string    `db:"ip_address"`
	Timezone     string    `db:"timezone"`
	Lat          float64   `db:"lat"`
	Lon          float64   `db:"lon"`
	AccessTime   time.Time `db:"access_time"`
}

type LoginLogRepository interface {
	FindLastAuthorized(ctx context.Context, userId uuid.UUID) (LoginLog, error)
	Save(ctx context.Context, login *LoginLog) error
}
