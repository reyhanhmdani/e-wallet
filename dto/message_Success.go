package dto

type SuccessCommonRes struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
}

type SuccessTNData struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Data    any    `json:"data"`
}
