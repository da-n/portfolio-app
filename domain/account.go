package domain

import (
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

const OrderSheetComplete = "complete"

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

type WithdrawalRequest struct {
	Id        int64  `db:"id"`
	AccountId int64  `db:"account_id"`
	Amount    int64  `db:"amount"`
	CreatedAt string `db:"created_at"`
}

// ToDto takes a WithdrawalRequest and casts it to dto.WithdrawalRequestResponse
func (w WithdrawalRequest) ToDto() dto.WithdrawalRequestResponse {
	return dto.WithdrawalRequestResponse{
		Id:        w.Id,
		AccountId: w.AccountId,
		Amount:    w.Amount,
		CreatedAt: w.CreatedAt,
	}
}

type OrderSheet struct {
	Id                  int64  `db:"id"`
	AccountId           int64  `db:"account_id"`
	WithdrawalRequestId int64  `db:"withdrawal_request_id"`
	Status              string `db:"status"`
	CreatedAt           string `db:"created_at"`
}

// ToDto takes a OrderSheet and casts it to dto.OrderSheetResponse
func (w OrderSheet) ToDto() dto.OrderSheetResponse {
	return dto.OrderSheetResponse{
		Id:                  w.Id,
		AccountId:           w.AccountId,
		WithdrawalRequestId: w.WithdrawalRequestId,
		Status:              w.Status,
		CreatedAt:           w.CreatedAt,
	}
}

type Instruction struct {
	Id              int64  `db:"id"`
	OrderSheetId    int64  `db:"customer_id"`
	InstructionType string `db:"instruction_type"`
	Amount          int64  `db:"amount"`
	CurrencyCode    string `db:"currency_code"`
}

// ToDto takes an Instruction and casts it to dto.InstructionResponse
func (a Instruction) ToDto() dto.InstructionResponse {
	return dto.InstructionResponse{
		Id:              a.Id,
		OrderSheetId:    a.OrderSheetId,
		InstructionType: a.InstructionType,
		Amount:          a.Amount,
		CurrencyCode:    a.CurrencyCode,
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
