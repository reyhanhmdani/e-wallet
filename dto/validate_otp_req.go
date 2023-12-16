package dto

type ValidateOTPReq struct {
	ReferenceID string `json:"reference_id"`
	OTP         string `json:"otp"`
}
