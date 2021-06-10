package service

import (
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type AccountService interface {
	ListAccounts(customerId int) ([]dto.AccountResponse, *errs.AppError)
	GetAccount(accountId int) (*dto.AccountResponse, *errs.AppError)
	CreateWithdrawalRequest(req *dto.WithdrawalRequestRequest) (*dto.WithdrawalRequestResponse, *errs.AppError)
	GetOrderSheet(orderSheetId int) (*dto.OrderSheetResponse, *errs.AppError)
	CreateOrderSheet(withdrawalRequest *domain.WithdrawalRequest) (*dto.OrderSheetResponse, *errs.AppError)
	CreateInstructions(orderSheet *domain.OrderSheet, withdrawalRequest *domain.WithdrawalRequest, account *domain.Account) ([]domain.Instruction, *errs.AppError)
	GetPortfolio(portfolioId int) (*dto.PortfolioResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
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

	a := domain.WithdrawalRequest{
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

func (service DefaultAccountService) CreateOrderSheet(withdrawalRequest *domain.WithdrawalRequest) (*dto.OrderSheetResponse, *errs.AppError) {
	// demo comments
	// for r the purposes of the demo the order sheet is going to be created amd completed synchronously, it could be
	// enhanced by initially returning it as "pending" and any consumer can poll or receive notification when it has
	// been "completed"
	o := domain.OrderSheet{
		WithdrawalRequestId: withdrawalRequest.Id,
		Status:              domain.OrderSheetComplete,
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

func (service DefaultAccountService) CreateInstructions(orderSheet *domain.OrderSheet, withdrawalRequest *domain.WithdrawalRequest, account *domain.Account) ([]domain.Instruction, *errs.AppError) {
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
	instructions := make([]domain.Instruction, 0)
	for _, asset := range portfolio.Assets {
		i := domain.Instruction{
			OrderSheetId:    orderSheet.Id,
			InstructionType: domain.InstructionTypeSell,
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

func NewAccountService(r domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{r}
}
