package dto

type OrderSheetResponse struct {
	Id                  int64  `json:"id"`
	AccountId           int64  `json:"accountId"`
	WithdrawalRequestId int64  `json:"withdrawalRequestId"`
	Status              string `json:"status"`
	CreatedAt           string `json:"createdAt"`
}
