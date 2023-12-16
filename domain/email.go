package domain

type EmailService interface {
	Send(to, subject, body string) error
	SendOTPByEmail(email, subject, body string) error
}
