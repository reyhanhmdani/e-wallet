package repository

import (
	"context"
	"database/sql"
	"e-wallet/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db *goqu.Database
}

func NewUser(con *sql.DB) domain.UserRepository {
	return &UserRepository{
		db: goqu.New("default", con),
	}
}

func (u UserRepository) FindByID(c context.Context, id uuid.UUID) (user domain.User, err error) {
	dataset := u.db.From("users").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(c, &user)
	return
}

func (u UserRepository) AllUsers(c context.Context) (users []domain.Users, err error) {
	dataset := u.db.From("users").Select("id", "username", "fullname", "email", "phone", "email_verified_at")
	err = dataset.ScanStructsContext(c, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// find user criteria
func (u UserRepository) FindUser(c context.Context, criteria map[string]interface{}, useAnd bool) (user domain.User, err error) {
	expressions := make([]goqu.Expression, 0)

	for key, value := range criteria {
		expressions = append(expressions, goqu.I(key).Eq(value))
	}

	dataset := u.db.From("users")

	if len(expressions) > 0 {
		if useAnd {
			dataset = dataset.Where(goqu.And(expressions...))
		} else {
			dataset = dataset.Where(goqu.Or(expressions...))
		}
	}

	_, err = dataset.ScanStructContext(c, &user)
	return
}

func (u UserRepository) Insert(c context.Context, user *domain.User) error {
	executor := u.db.Insert("users").Rows(goqu.Record{
		"fullname": user.Fullname,
		"username": user.Username,
		"password": user.Password,
		"phone":    user.Phone,
		"email":    user.Email,
	}).Returning("id").Executor() // supaya mendapatkan id nya
	_, err := executor.ScanStructContext(c, user)
	return err
}

func (u UserRepository) Update(c context.Context, user *domain.User) error {
	user.EmailVerifiedAtDB = sql.NullTime{
		Time:  user.EmailVerifiedAt,
		Valid: true,
	}
	executor := u.db.Update("users").Where(goqu.Ex{
		"id": user.ID,
	}).Set(goqu.Record{
		"fullname":          user.Fullname,
		"username":          user.Username,
		"password":          user.Password,
		"phone":             user.Phone,
		"email":             user.Email,
		"email_verified_at": user.EmailVerifiedAtDB,
	}).Executor() // supaya mendapatkan id nya
	sqltest, _, _ := executor.ToSQL()
	logrus.Info("SQL:", sqltest)
	_, err := executor.ExecContext(c)
	//logrus.Error("ga ada error", err)
	return err
}

// dibawah ini sudah tidak terpakai lagi,find find nya di karenakan tergantikan dengan FindUser
func (u UserRepository) FindByUsernameOrEmail(c context.Context, username, email string) (user domain.User, err error) {
	datasetOr := u.db.From("users").Where(
		goqu.And(
			goqu.Ex{"username": username},
			goqu.Ex{"email": email},
		),
	)

	_, err = datasetOr.ScanStructContext(c, &user)
	return
}
func (u UserRepository) FindByUsername(c context.Context, username string) (user domain.User, err error) {
	datasetOr := u.db.From("users").Where(
		goqu.Ex{"username": username},
	)

	_, err = datasetOr.ScanStructContext(c, &user)
	return
}
func (u UserRepository) FindByEmail(c context.Context, email string) (user domain.User, err error) {
	datasetOr := u.db.From("users").Where(
		goqu.Ex{"email": email},
	)

	_, err = datasetOr.ScanStructContext(c, &user)
	return
}
