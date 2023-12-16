package dto

import (
	"github.com/google/uuid"
	"time"
)

type NotificationData struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"` // menampung isi dari notification nya
	Status    int8      `json:"status"`
	IsRead    int8      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
