package dto

type EmailReq struct {
	Email string `json:"email"`
}
type EmailSendReq struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
