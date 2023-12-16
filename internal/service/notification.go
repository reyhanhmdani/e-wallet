package service

import (
	"bytes"
	"context"
	"e-wallet/domain"
	"e-wallet/dto"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"html/template"
	"time"
)

type NotificationService struct {
	notificationRepository domain.NotificationRepository
	templateRepository     domain.TemplateRepository
	hub                    *dto.Hub
}

func NewNotification(notificationRepoitory domain.NotificationRepository,
	templateRepository domain.TemplateRepository,
	hub *dto.Hub) domain.NotificationService {
	return &NotificationService{
		notificationRepository: notificationRepoitory,
		templateRepository:     templateRepository,
		hub:                    hub,
	}
}

func (n NotificationService) FindByUser(c context.Context, user uuid.UUID) ([]dto.NotificationData, error) {
	notification, err := n.notificationRepository.FindByUser(c, user)
	if err != nil {
		logrus.Error("error di service", err)
		return nil, err
	}
	var result []dto.NotificationData
	for _, v := range notification {
		result = append(result, dto.NotificationData{
			Id:        v.Id,
			Title:     v.Title,
			Body:      v.Body,
			Status:    v.Status,
			IsRead:    v.IsRead,
			CreatedAt: v.CreatedAt,
		})
	}
	if result == nil {
		// empty array
		result = make([]dto.NotificationData, 0)
	}

	return result, nil
}

func (n NotificationService) Insert(c context.Context, userId uuid.UUID, code string, data map[string]string) error {
	tmpl, err := n.templateRepository.FindByCode(c, code)
	if err != nil {
		logrus.Error("Code tidak di temukan", err)
		return err
	}

	if tmpl == (domain.Template{}) {
		return domain.TemplNotFound
	}

	// untuk menerima hasil parsingannya
	body := new(bytes.Buffer)
	t := template.Must(template.New("notif").Parse(tmpl.Body))
	err = t.Execute(body, data)
	if err != nil {
		logrus.Error("error di bagain setelah parsing", err)
		return err
	}

	notification := domain.Notification{
		UserId:    userId,
		Title:     tmpl.Title,
		Body:      body.String(),
		Status:    1,
		IsRead:    0,
		CreatedAt: time.Now(),
	}

	err = n.notificationRepository.Insert(c, &notification)
	if err != nil {
		logrus.Error("error di bagian insert", err)
		return err
	}

	if channel, ok := n.hub.NotificationChannel[userId]; ok {
		channel <- dto.NotificationData{
			Id:        notification.Id,
			Title:     notification.Title,
			Body:      notification.Body,
			Status:    notification.Status,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		}
	}

	return nil
}
