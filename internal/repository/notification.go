package repository

import (
	"context"
	"database/sql"
	"e-wallet/domain"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// implementasi dari interface notification

type notificationRepository struct {
	db *goqu.Database
}

func NewNotification(conn *sql.DB) domain.NotificationRepository {
	return &notificationRepository{
		db: goqu.New("default", conn),
	}
}

func (n notificationRepository) FindByUser(c context.Context, user uuid.UUID) (notification []domain.Notification, err error) {
	dataset := n.db.From("notifications").Where(goqu.Ex{
		"user_id": user,
	}).Order(goqu.I("created_at").Desc()).Limit(15)
	err = dataset.ScanStructsContext(c, &notification)
	return
}

func (n notificationRepository) Insert(c context.Context, notification *domain.Notification) error {
	executor := n.db.Insert("notifications").Rows(goqu.Record{
		"user_id":    notification.UserId,
		"title":      notification.Title,
		"body":       notification.Body,
		"status":     notification.Status,
		"is_read":    notification.IsRead,
		"created_at": notification.CreatedAt,
	}).Returning("id").Executor()
	_, err := executor.ScanStructContext(c, notification)
	if err != nil {
		logrus.Error("Insert notifications failed: ", err)
		return err
	}
	return err
}

func (n notificationRepository) Update(c context.Context, notification *domain.Notification) error {
	executor := n.db.Update("notifications").Set(goqu.Ex{
		"id": notification.Id,
	}).Set(goqu.Record{
		"title":   notification.Title,
		"body":    notification.Body,
		"status":  notification.Status,
		"is_read": notification.IsRead,
	}).Executor()
	_, err := executor.ExecContext(c)
	if err != nil {
		logrus.Error("Update notification failed: ", err)
		return err
	}
	return err
}
