package domain

import (
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	AccountType string `db:"account_type"`
	Balance     int    `db:"balance"`
	OpeningDate string `db:"opening_date"`
	Status      string `db:"status"`
}

// ToDto takes a Account and casts it to dto.AccountResponse
func (a Account) ToDto() dto.AccountResponse {
	return dto.AccountResponse{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerId,
		AccountType: a.AccountType,
		Balance:     a.Balance,
		OpeningDate: a.OpeningDate,
		Status:      a.Status,
	}
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/da-n/portfolio-app/domain AccountRepository
type AccountRepository interface {
	FindAll() ([]Account, *errs.AppError)
	FindByAccountId(string) (*Account, *errs.AppError)
}
