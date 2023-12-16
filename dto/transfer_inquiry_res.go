package dto

type TransferInquiryRes struct {
	// key ini yang di pakai ketika satu proses yang lain yaitu yang namanya transferExecuteReq
	InquiryKey string `json:"inquiry_key"`
}
