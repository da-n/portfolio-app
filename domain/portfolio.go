package domain

import (
	"github.com/da-n/portfolio-app/dto"
)

type Portfolio struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

// ToDto takes a Portfolio and casts it to dto.PortfolioResponse
func (a Portfolio) ToDto() dto.PortfolioResponse {
	return dto.PortfolioResponse{
		Id:   a.Id,
		Name: a.Name,
	}
}