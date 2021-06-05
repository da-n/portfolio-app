package domain

import (
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type Customer struct {
	CustomerId string `db:"id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Email      string `db:"email"`
	Password   string `db:"password"`
}

// ToDto takes a Customer and casts it to dto.CustomerResponse
func (u Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		CustomerId: u.CustomerId,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Email:      u.Email,
	}
}

//go:generate mockgen -destination=../mocks/domain/mockCustomerRepository.go -package=domain github.com/da-n/portfolio-app/domain CustomerRepository
type CustomerRepository interface {
	FindById(customerId string) (*Customer, *errs.AppError)
}
