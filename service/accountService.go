package service

import (
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	ListAccounts(customerId int64) ([]dto.AccountResponse, *errs.AppError)
	GetAccount(accountId int64) (*dto.AccountResponse, *errs.AppError)
	CreateWithdrawalRequest(req dto.WithdrawalRequestRequest) (*dto.WithdrawalRequestResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (service DefaultAccountService) ListAccounts(customerId int64) ([]dto.AccountResponse, *errs.AppError) {
	a, err := service.repo.FindAll(customerId)
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
	a, err := service.repo.FindById(accountId)
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
	account, err := service.repo.SaveWithdrawalRequest(a)
	if err != nil {
		return nil, err
	}

	response := account.ToDto()
	return &response, nil
}

func NewAccountService(r domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{r}
}
