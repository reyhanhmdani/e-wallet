package dto

import "github.com/google/uuid"

type UserData struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Username string    `json:"username"`
}
