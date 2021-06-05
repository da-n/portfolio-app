package domain

import (
	"database/sql"
	"github.com/da-n/portfolio-app/errs"
	"github.com/da-n/portfolio-app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// CustomerRepositoryDb is the database implementation of CustomerRepository
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

// FindById find a customer by their customer_id
func (r CustomerRepositoryDb) FindByCustomerId(customerId string) (*Customer, *errs.AppError) {
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
