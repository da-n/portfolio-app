package dto

type AccountResponse struct {
	Id           int    `json:"id"`
	CustomerId   int    `json:"customerId"`
	PortfolioId  int    `json:"portfolioId"`
	CurrencyCode string `json:"currencyCode"`
	Balance      int    `json:"balance"`
}
