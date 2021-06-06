package domain

import (
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type Account struct {
	Id          int64 `db:"id"`
	CustomerId  int64 `db:"customer_id"`
	PortfolioId int64 `db:"portfolio_id"`
	Balance     int64 `db:"balance"`
}

// ToDto takes a Account and casts it to dto.AccountResponse
func (a Account) ToDto() dto.AccountResponse {
	return dto.AccountResponse{
		Id:          a.Id,
		CustomerId:  a.CustomerId,
		PortfolioId: a.PortfolioId,
		Balance:     a.Balance,
	}
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/da-n/portfolio-app/domain AccountRepository
type AccountRepository interface {
	FindAllAccounts(int64) ([]Account, *errs.AppError)
	FindAccountById(int64) (*Account, *errs.AppError)
	SaveWithdrawalRequest(WithdrawalRequest) (*WithdrawalRequest, *errs.AppError)
	FindOrderSheetById(int64) (*OrderSheet, *errs.AppError)
	SaveOrderSheet(OrderSheet) (*OrderSheet, *errs.AppError)
}
