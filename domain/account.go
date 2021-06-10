package domain

import (
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

const OrderSheetComplete = "complete"
const InstructionTypeSell = "SELL"

type Account struct {
	Id           int    `db:"id"`
	CustomerId   int    `db:"customer_id"`
	PortfolioId  int    `db:"portfolio_id"`
	CurrencyCode string `db:"currency_code"`
	Balance      int    `db:"balance"`
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
	Id        int    `db:"id"`
	AccountId int    `db:"account_id"`
	Amount    int    `db:"amount"`
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
	Id                  int    `db:"id"`
	WithdrawalRequestId int    `db:"withdrawal_request_id"`
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
	Id              int    `db:"id"`
	OrderSheetId    int    `db:"order_sheet_id"`
	InstructionType string `db:"instruction_type"`
	Isin            string `db:"isin"`
	Amount          int    `db:"amount"`
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
	Id     int     `db:"id"`
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
	Id      int    `db:"id"`
	Isin    string `db:"isin"`
	Name    string `db:"name"`
	Percent int    `db:"percent"`
}

func (a Asset) ToDto() dto.AssetResponse {
	return dto.AssetResponse{
		Id:      a.Id,
		Isin:    a.Isin,
		Name:    a.Name,
		Percent: a.Percent,
	}
}

//go:generate mockgen -destination=../domain/account_mock.go -package=domain github.com/da-n/portfolio-app/domain AccountRepository
type AccountRepository interface {
	FindAllAccounts(customerId int) ([]Account, *errs.AppError)
	FindAccountById(accountId int) (*Account, *errs.AppError)
	SaveWithdrawalRequest(withdrawalRequest WithdrawalRequest) (*WithdrawalRequest, *errs.AppError)
	FindOrderSheetById(orderSheetId int) (*OrderSheet, *errs.AppError)
	SaveOrderSheet(orderSheet OrderSheet) (*OrderSheet, *errs.AppError)
	FindPortfolioById(portfolioId int) (*Portfolio, *errs.AppError)
	SaveInstruction(instruction Instruction) (*Instruction, *errs.AppError)
}

type AccountService interface {
	ListAccounts(customerId int) ([]dto.AccountResponse, *errs.AppError)
	GetAccount(accountId int) (*dto.AccountResponse, *errs.AppError)
	CreateWithdrawalRequest(req *dto.WithdrawalRequestRequest) (*dto.WithdrawalRequestResponse, *errs.AppError)
	GetOrderSheet(orderSheetId int) (*dto.OrderSheetResponse, *errs.AppError)
	CreateOrderSheet(withdrawalRequest *WithdrawalRequest) (*dto.OrderSheetResponse, *errs.AppError)
	CreateInstructions(orderSheet *OrderSheet, withdrawalRequest *WithdrawalRequest, account *Account) ([]Instruction, *errs.AppError)
	GetPortfolio(portfolioId int) (*dto.PortfolioResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo AccountRepository
}

func (service DefaultAccountService) ListAccounts(customerId int) ([]dto.AccountResponse, *errs.AppError) {
	a, err := service.repo.FindAllAccounts(customerId)
	if err != nil {
		return nil, err
	}

	response := make([]dto.AccountResponse, 0)
	for _, account := range a {
		response = append(response, account.ToDto())
	}

	return response, nil
}

func (service DefaultAccountService) GetAccount(accountId int) (*dto.AccountResponse, *errs.AppError) {
	a, err := service.repo.FindAccountById(accountId)
	if err != nil {
		return nil, err
	}

	response := a.ToDto()

	return &response, nil
}

func (service DefaultAccountService) CreateWithdrawalRequest(req *dto.WithdrawalRequestRequest) (*dto.WithdrawalRequestResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := WithdrawalRequest{
		AccountId: req.AccountId,
		Amount:    req.Amount,
	}
	withdrawalRequest, err := service.repo.SaveWithdrawalRequest(a)
	if err != nil {
		return nil, err
	}

	orderSheet, err := service.CreateOrderSheet(withdrawalRequest)
	if err != nil {
		return nil, err
	}

	response := withdrawalRequest.ToDto()
	response.OrderSheet = orderSheet

	return &response, nil
}

func (service DefaultAccountService) GetOrderSheet(orderSheetId int) (*dto.OrderSheetResponse, *errs.AppError) {
	o, err := service.repo.FindOrderSheetById(orderSheetId)
	if err != nil {
		return nil, err
	}

	response := o.ToDto()

	return &response, nil
}

func (service DefaultAccountService) CreateOrderSheet(withdrawalRequest *WithdrawalRequest) (*dto.OrderSheetResponse, *errs.AppError) {
	// demo comments
	// for r the purposes of the demo the order sheet is going to be created amd completed synchronously, it could be
	// enhanced by initially returning it as "pending" and any consumer can poll or receive notification when it has
	// been "completed"
	o := OrderSheet{
		WithdrawalRequestId: withdrawalRequest.Id,
		Status:              OrderSheetComplete,
	}

	orderSheet, err := service.repo.SaveOrderSheet(o)
	if err != nil {
		return nil, err
	}

	account, err := service.repo.FindAccountById(withdrawalRequest.AccountId)
	if err != nil {
		return nil, err
	}

	instructions, err := service.CreateInstructions(orderSheet, withdrawalRequest, account)
	if err != nil {
		return nil, err
	}
	orderSheet.Instructions = instructions
	response := orderSheet.ToDto()

	return &response, nil
}

func (service DefaultAccountService) CreateInstructions(orderSheet *OrderSheet, withdrawalRequest *WithdrawalRequest, account *Account) ([]Instruction, *errs.AppError) {
	portfolio, err := service.repo.FindPortfolioById(account.PortfolioId)
	if err != nil {
		return nil, err
	}

	// demo comments
	// we are going to divide the total amount requested by the percentages of the portfolio assets holdings e.g.
	// asset a = 40%
	// asset b = 60%
	// amount requested = £1000
	// instruction 1 = sell 400 for asset a
	// instruction 2 = sell 600 for asset b
	// using basic formula: percentage of asset * amount requested / 100
	instructions := make([]Instruction, 0)
	for _, asset := range portfolio.Assets {
		i := Instruction{
			OrderSheetId:    orderSheet.Id,
			InstructionType: InstructionTypeSell,
			Isin:            asset.Isin,
			Amount:          (asset.Percent * withdrawalRequest.Amount) / 100,
			CurrencyCode:    account.CurrencyCode,
		}
		instruction, err := service.repo.SaveInstruction(i)
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, *instruction)
	}

	return instructions, nil
}

func (service DefaultAccountService) GetPortfolio(portfolioId int) (*dto.PortfolioResponse, *errs.AppError) {
	p, err := service.repo.FindPortfolioById(portfolioId)
	if err != nil {
		return nil, err
	}

	response := p.ToDto()

	return &response, nil
}

func NewAccountService(r AccountRepository) DefaultAccountService {
	return DefaultAccountService{r}
}