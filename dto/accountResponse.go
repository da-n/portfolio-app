package dto

type AccountResponse struct {
	Id           int64  `json:"id"`
	CustomerId   int64  `json:"customerId"`
	PortfolioId  int64  `json:"portfolioId"`
	CurrencyCode string `json:"currencyCode"`
	Balance      int64  `json:"balance"`
}
