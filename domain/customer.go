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

//go:generate mockgen -destination=../mocks/domain/mockCustomerRepository.go -package=domain github.com/da-n/portfolio-app/domain CustomerRepository
type CustomerRepository interface {
	FindByCustomerId(customerId int) (*Customer, *errs.AppError)
}
