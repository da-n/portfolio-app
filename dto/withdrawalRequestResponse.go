package dto

type WithdrawalRequestResponse struct {
	Id         int64              `json:"id"`
	AccountId  int64              `json:"accountId"`
	Amount     int64              `json:"amount"`
	CreatedAt  string             `json:"createdAt"`
	OrderSheet OrderSheetResponse `json:"orderSheet"`
}
