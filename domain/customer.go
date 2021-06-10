package domain

import (
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type Customer struct {
	Id        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Language  string `db:"language"`
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:        c.Id,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
		Language:  c.Language,
	}
}

//go:generate mockgen -destination=../domain/customer_mock.go -package=domain github.com/da-n/portfolio-app/domain CustomerRepository
type CustomerRepository interface {
	FindByCustomerId(customerId int) (*Customer, *errs.AppError)
}

type CustomerService interface {
	GetCustomer(customerId int) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo CustomerRepository
}

func (service DefaultCustomerService) GetCustomer(customerId int) (*dto.CustomerResponse, *errs.AppError) {
	c, err := service.repo.FindByCustomerId(customerId)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(r CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{r}
}