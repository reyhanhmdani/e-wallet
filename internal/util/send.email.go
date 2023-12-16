package util

import (
	"e-wallet/internal/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendOTPByEmail(email, otp string) error {
	cnf := config.Get().Email
	m := gomail.NewMessage()
	m.SetHeader("From", cnf.EmailFrom)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Kode OTP Anda")
	m.SetBody("text/html", "ini adalah kode otp anda "+otp)

	port, err := strconv.Atoi(cnf.Port)
	if err != nil {
		logrus.Error("Error converting port to integer: ", err)
		return err
	}

	d := gomail.NewDialer(cnf.Host, port, cnf.User, cnf.Password)

	// Kirim email
	if err := d.DialAndSend(m); err != nil {
		logrus.Error("error di Dial And Send ", err)
		return err
	}

	return nil
}
