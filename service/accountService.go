package service

import (
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type AccountService interface {
	ListAccounts() ([]dto.AccountResponse, *errs.AppError)
	GetAccount(string) (*dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (service DefaultAccountService) ListAccounts() ([]dto.AccountResponse, *errs.AppError) {
	a, err := service.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.AccountResponse, 0)
	for _, account := range a {
		response = append(response, account.ToDto())
	}

	return response, nil
}

func (service DefaultAccountService) GetAccount(accountId string) (*dto.AccountResponse, *errs.AppError) {
	a, err := service.repo.FindById(accountId)
	if err != nil {
		return nil, err
	}

	response := a.ToDto()
	return &response, nil
}

func NewAccountService(r domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{r}
}
