package domain

import (
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

const OrderSheetComplete = "complete"
const InstructionTypeSell = "SELL"

type Account struct {
	Id           int64  `db:"id"`
	CustomerId   int64  `db:"customer_id"`
	PortfolioId  int64  `db:"portfolio_id"`
	CurrencyCode string `db:"currency_code"`
	Balance      int64  `db:"balance"`
}

func (a Account) ToDto() dto.AccountResponse {
	return dto.AccountResponse{
		Id:           a.Id,
		CustomerId:   a.CustomerId,
		PortfolioId:  a.PortfolioId,
		CurrencyCode: a.CurrencyCode,
		Balance:      a.Balance,
	}
}

type WithdrawalRequest struct {
	Id        int64  `db:"id"`
	AccountId int64  `db:"account_id"`
	Amount    int64  `db:"amount"`
	CreatedAt string `db:"created_at"`
}

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
	WithdrawalRequestId int64  `db:"withdrawal_request_id"`
	Status              string `db:"status"`
	CreatedAt           string `db:"created_at"`
	Instructions        []Instruction
}

func (o OrderSheet) ToDto() dto.OrderSheetResponse {
	instructions := make([]dto.InstructionResponse, 0)
	for _, v := range o.Instructions {
		instructions = append(instructions, v.ToDto())
	}

	return dto.OrderSheetResponse{
		Id:                  o.Id,
		WithdrawalRequestId: o.WithdrawalRequestId,
		Status:              o.Status,
		CreatedAt:           o.CreatedAt,
		Instructions:        &instructions,
	}
}

type Instruction struct {
	Id              int64  `db:"id"`
	OrderSheetId    int64  `db:"order_sheet_id"`
	InstructionType string `db:"instruction_type"`
	Isin            string `db:"isin"`
	Amount          int64  `db:"amount"`
	CurrencyCode    string `db:"currency_code"`
	CreatedAt       string `db:"created_at"`
}

func (i Instruction) ToDto() dto.InstructionResponse {
	return dto.InstructionResponse{
		Id:              i.Id,
		OrderSheetId:    i.OrderSheetId,
		InstructionType: i.InstructionType,
		Isin:            i.Isin,
		Amount:          i.Amount,
		CurrencyCode:    i.CurrencyCode,
		CreatedAt:       i.CreatedAt,
	}
}

type Portfolio struct {
	Id     int64   `db:"id"`
	Name   string  `db:"name"`
	Assets []Asset `db:"assets"`
}

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

func (a Asset) ToDto() dto.AssetResponse {
	return dto.AssetResponse{
		Id:      a.Id,
		Isin:    a.Isin,
		Name:    a.Name,
		Percent: a.Percent,
	}
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/da-n/portfolio-app/domain AccountRepository
type AccountRepository interface {
	FindAllAccounts(customerId int64) ([]Account, *errs.AppError)
	FindAccountById(accountId int64) (*Account, *errs.AppError)
	SaveWithdrawalRequest(withdrawalRequest WithdrawalRequest) (*WithdrawalRequest, *errs.AppError)
	FindOrderSheetById(orderSheetId int64) (*OrderSheet, *errs.AppError)
	SaveOrderSheet(orderSheet OrderSheet) (*OrderSheet, *errs.AppError)
	FindPortfolioById(portfolioId int64) (*Portfolio, *errs.AppError)
	SaveInstruction(instruction Instruction) (*Instruction, *errs.AppError)
}
