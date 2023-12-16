package service

import (
	"e-wallet/domain"
	"e-wallet/dto"
	"e-wallet/internal/config"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"strconv"
)

type EmailService struct {
	cnf          *config.Config
	queueService domain.QueueService
}

func NewEmail(cnf *config.Config, queueService domain.QueueService) domain.EmailService {
	return &EmailService{
		cnf,
		queueService,
	}
}

func (e EmailService) Send(to, subject, body string) error {

	payload := dto.EmailSendReq{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return e.queueService.Enqueue("send:email", data, 3)
}

func (e EmailService) SendOTPByEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.cnf.Email.EmailFrom)
	m.SetHeader("To", to)
	//m.SetHeader("Subject", "Kode OTP Anda")
	//m.SetBody("text/html", "ini adalah kode otp anda "+otp)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// 1. Konversi data ke struct
	data := dto.EmailSendReq{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	// 2. Marshal struct menjadi byte array
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// 3. Panggil queueService.Enqueue
	err = e.queueService.Enqueue("send:email", payload, 3)
	if err != nil {
		return err
	}

	// 4. Kirim email asynchronously
	go func() {
		port, err := strconv.Atoi(e.cnf.Email.Port)
		if err != nil {
			logrus.Error("Error converting port to integer: ", err)
			return
		}
		d := gomail.NewDialer(e.cnf.Email.Host, port, e.cnf.Email.User, e.cnf.Email.Password)
		if err := d.DialAndSend(m); err != nil {
			logrus.Error("error di Dial And Send ", err)
			return
		}
	}()

	return nil
}
