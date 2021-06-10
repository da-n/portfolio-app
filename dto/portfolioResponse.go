package dto

type PortfolioResponse struct {
	Id     int              `json:"id"`
	Name   string           `json:"name"`
	Assets *[]AssetResponse `json:"assets"`
}
