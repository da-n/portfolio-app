package dto

type WithdrawalRequestResponse struct {
	Id         int                 `json:"id"`
	AccountId  int                 `json:"accountId"`
	Amount     int                 `json:"amount"`
	CreatedAt  string              `json:"createdAt"`
	OrderSheet *OrderSheetResponse `json:"orderSheet"`
}
