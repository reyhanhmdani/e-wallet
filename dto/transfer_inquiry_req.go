package dto

type TransferInquiryReq struct {
	// account number tujuan
	AccountNumber string `json:"account_number"`
	// mau ngirim berapa
	Amount float64 `json:"amount"`
}
