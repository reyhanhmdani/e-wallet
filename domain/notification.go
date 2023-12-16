package domain

import (
	"context"
	"e-wallet/dto"
	"github.com/google/uuid"
	"time"
)

type Notification struct {
	Id        uuid.UUID `db:"id"`
	UserId    uuid.UUID `db:"user_id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"` // menampung isi dari notification nya
	Status    int8      `db:"status"`
	IsRead    int8      `db:"is_read"`
	CreatedAt time.Time `db:"created_at"`
}

type NotificationRepository interface {
	FindByUser(c context.Context, user uuid.UUID) ([]Notification, error)
	Insert(c context.Context, notification *Notification) error
	Update(c context.Context, notification *Notification) error
}

type NotificationService interface {
	FindByUser(c context.Context, user uuid.UUID) ([]dto.NotificationData, error)
	Insert(c context.Context, userId uuid.UUID, code string, data map[string]string) error
}
