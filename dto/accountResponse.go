package dto

type AccountResponse struct {
	AccountId   string `json:"account_id"`
	CustomerId  string `json:"customer_id"`
	AccountType string `json:"account_type"`
	Balance     string `json:"balance"`
}
