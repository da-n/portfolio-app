package service

import (
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type CustomerService interface {
	GetCustomer(customerId int64) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (service DefaultCustomerService) GetCustomer(customerId int64) (*dto.CustomerResponse, *errs.AppError) {
	c, err := service.repo.FindById(customerId)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(r domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{r}
}
