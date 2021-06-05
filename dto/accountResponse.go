package dto

type AccountResponse struct {
	Id          string `json:"id"`
	CustomerId  string `json:"customerId"`
	AccountType string `json:"accountType"`
	Balance     int    `json:"balance"`
}
