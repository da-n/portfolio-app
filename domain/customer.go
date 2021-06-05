package domain

import (
	"database/sql"
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
	"github.com/da-n/portfolio-app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Customer struct {
	CustomerId string `db:"customer_id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Email      string `db:"email"`
	Password   string `db:"password"`
}

func (u Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		CustomerId: u.CustomerId,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		Email:      u.Email,
	}
}

type CustomerRepository interface {
	FindById(string) (*Customer, *errs.AppError)
}

// CustomerRepositoryDb is the database implementation of the repository
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (r CustomerRepositoryDb) FindById(customerId string) (*Customer, *errs.AppError) {
	query := "select customer_id, first_name, last_name, email from customers where customers.customer_id = ?"

	var c Customer
	err := r.client.Get(&c, query, customerId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer could not be found")
		} else {
			logger.Error("Error while querying customers: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

// NewCustomerRepositoryDb instantiates a new CustomerRepositoryDb passing in a sqlx.DB instance
func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
