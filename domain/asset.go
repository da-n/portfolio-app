package domain

import (
	"github.com/da-n/portfolio-app/dto"
)

type Asset struct {
	Id   int64  `db:"id"`
	Isin string `db:"isin"`
	Name string `db:"name"`
}

// ToDto takes a Account and casts it to dto.AccountResponse
func (a Asset) ToDto() dto.AssetResponse {
	return dto.AssetResponse{
		Id:   a.Id,
		Isin: a.Isin,
		Name: a.Name,
	}
}
