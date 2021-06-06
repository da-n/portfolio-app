package service

import (
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	ListAccounts(int64) ([]dto.AccountResponse, *errs.AppError)
	GetAccount(int64) (*dto.AccountResponse, *errs.AppError)
	CreateWithdrawalRequest(dto.WithdrawalRequestRequest) (*dto.WithdrawalRequestResponse, *errs.AppError)
	GetOrderSheet(int64) (*dto.OrderSheetResponse, *errs.AppError)
	CreateOrderSheet(domain.WithdrawalRequest) (*domain.OrderSheet, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (service DefaultAccountService) ListAccounts(customerId int64) ([]dto.AccountResponse, *errs.AppError) {
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

func (service DefaultAccountService) GetAccount(accountId int64) (*dto.AccountResponse, *errs.AppError) {
	a, err := service.repo.FindAccountById(accountId)
	if err != nil {
		return nil, err
	}

	response := a.ToDto()
	return &response, nil
}

func (service DefaultAccountService) CreateWithdrawalRequest(req dto.WithdrawalRequestRequest) (*dto.WithdrawalRequestResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	a := domain.WithdrawalRequest{
		AccountId: req.AccountId,
		Amount:    req.Amount,
		CreatedAt: time.Now().Format(dbTSLayout),
	}
	withdrawalRequest, err := service.repo.SaveWithdrawalRequest(a)
	if err != nil {
		return nil, err
	}

	orderSheet, err := service.CreateOrderSheet(*withdrawalRequest)
	if err != nil {
		return nil, err
	}

	response := withdrawalRequest.ToDto()
	response.OrderSheet = orderSheet.ToDto()
	return &response, nil
}

func (service DefaultAccountService) GetOrderSheet(orderSheetId int64) (*dto.OrderSheetResponse, *errs.AppError) {
	o, err := service.repo.FindOrderSheetById(orderSheetId)
	if err != nil {
		return nil, err
	}

	response := o.ToDto()
	return &response, nil
}

func (service DefaultAccountService) CreateOrderSheet(w domain.WithdrawalRequest) (*domain.OrderSheet, *errs.AppError) {

	// For the purposes of the demo the order sheet is going to be created amd completed synchronously, it would be
	// enhanced by initially returning it as "pending" and any consumer can poll or receive notification when it has
	// been "completed"
	o := domain.OrderSheet{
		AccountId:           w.AccountId,
		WithdrawalRequestId: w.Id,
		Status:              domain.OrderSheetComplete,
		CreatedAt:           time.Now().Format(dbTSLayout),
	}
	orderSheet, err := service.repo.SaveOrderSheet(o)
	if err != nil {
		return nil, err
	}

	// todo: implement instruction logic when entities and portfolios exist

	return orderSheet, nil
}

func NewAccountService(r domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{r}
}
