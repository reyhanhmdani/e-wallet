package domain

import (
	"context"
	"database/sql"
	"e-wallet/dto"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" db:"id"`
	Fullname          string       `db:"fullname"`
	Phone             string       `db:"phone"`
	Email             string       `db:"email"`
	Username          string       `db:"username"`
	Password          string       `db:"password"`
	EmailVerifiedAtDB sql.NullTime `db:"email_verified_at"`
	EmailVerifiedAt   time.Time    `db:"-"`
}

type Users struct {
	ID                uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" db:"id"`
	Fullname          string       `db:"fullname"`
	Phone             string       `db:"phone"`
	Email             string       `db:"email"`
	Username          string       `db:"username"`
	EmailVerifiedAtDB sql.NullTime `db:"email_verified_at"`
}

type UserRepository interface {
	FindByID(c context.Context, id uuid.UUID) (User, error)
	AllUsers(c context.Context) ([]Users, error)
	// Find criteria
	FindUser(c context.Context, criteria map[string]interface{}, useAndOr bool) (user User, err error)
	// sudah tergantikan dengan FindUser
	FindByUsernameOrEmail(c context.Context, username, email string) (User, error)
	FindByUsername(c context.Context, username string) (User, error)
	FindByEmail(c context.Context, email string) (User, error)

	Insert(c context.Context, user *User) error
	Update(c context.Context, user *User) error
}

type UserService interface {
	AllUsers(c context.Context) ([]Users, error)
	Authenticate(c context.Context, req dto.AuthReq) (dto.AuthRes, error)
	ValidateTokenWithCache(c context.Context, token string) (dto.UserData, error)
	ValidateTokenJwt(c context.Context, token string) (dto.UserData, error)
	Register(c context.Context, req dto.UserCreateReq) (dto.UserRegisterRes, error)
	ValidateOTP(c context.Context, req dto.ValidateOTPReq) error
}
