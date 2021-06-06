package domain

import (
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type Portfolio struct {
	Id     int64   `db:"id"`
	Name   string  `db:"name"`
	Assets []Asset `db:"assets"`
}

// ToDto takes a Portfolio and casts it to dto.PortfolioResponse
func (a Portfolio) ToDto() dto.PortfolioResponse {
	assets := make([]dto.AssetResponse, 0)
	for _, v := range a.Assets {
		assets = append(assets, v.ToDto())
	}
	return dto.PortfolioResponse{
		Id:     a.Id,
		Name:   a.Name,
		Assets: &assets,
	}
}

type Asset struct {
	Id      int64  `db:"id"`
	Isin    string `db:"isin"`
	Name    string `db:"name"`
	Percent int64  `db:"percent"`
}

// ToDto takes a Account and casts it to dto.AccountResponse
func (a Asset) ToDto() dto.AssetResponse {
	return dto.AssetResponse{
		Id:      a.Id,
		Isin:    a.Isin,
		Name:    a.Name,
		Percent: a.Percent,
	}
}

//go:generate mockgen -destination=../mocks/domain/mockPortfolioRepository.go -package=domain github.com/da-n/portfolio-app/domain PortfolioRepository
type PortfolioRepository interface {
	FindPortfolioById(int64) (*Portfolio, *errs.AppError)
}
