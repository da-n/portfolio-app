package dto

type AccountResponse struct {
	Id          int64  `json:"id"`
	CustomerId  int64  `json:"customerId"`
	AccountType string `json:"accountType"`
	Balance     int64  `json:"balance"`
}
