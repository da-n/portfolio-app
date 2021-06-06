package dto

type PortfolioResponse struct {
	Id     int64           `json:"id"`
	Name   string          `json:"name"`
	Assets []AssetResponse `json:"assets"`
}
