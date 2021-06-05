package dto

type AccountResponse struct {
	AccountId   string `json:"accountId"`
	CustomerId  string `json:"customerId"`
	AccountType string `json:"accountType"`
	Balance     int    `json:"balance"`
}
