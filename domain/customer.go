package domain

import (
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type Customer struct {
	Id        int64  `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
}

// ToDto takes a Customer and casts it to dto.CustomerResponse
func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:        c.Id,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Email:     c.Email,
	}
}

//go:generate mockgen -destination=../mocks/domain/mockCustomerRepository.go -package=domain github.com/da-n/portfolio-app/domain CustomerRepository
type CustomerRepository interface {
	FindById(customerId int64) (*Customer, *errs.AppError)
}
