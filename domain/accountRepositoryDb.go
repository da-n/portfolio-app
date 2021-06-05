package domain

import (
	"database/sql"
	"github.com/da-n/portfolio-app/errs"
	"github.com/da-n/portfolio-app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// AccountRepositoryDb is the database implementation of AccountRepository
type AccountRepositoryDb struct {
	client *sqlx.DB
}

// FindById find an account by its account_id
func (r AccountRepositoryDb) FindAll(customerId string) ([]Account, *errs.AppError) {
	query := "select id, customer_id, account_type, balance from accounts where customer_id = ?"
	a := make([]Account, 0)
	err := r.client.Select(&a, query, customerId)
	if err != nil {
		logger.Error("Error while querying accounts: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return a, nil
}

// FindById find an account by its account_id
func (r AccountRepositoryDb) FindById(accountId string) (*Account, *errs.AppError) {
	query := "select id, customer_id, account_type, balance from accounts where id = ?"

	var a Account
	err := r.client.Get(&a, query, accountId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account could not be found")
		} else {
			logger.Error("Error while querying accounts: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &a, nil
}

// NewAccountRepositoryDb instantiates a new AccountRepositoryDb passing in a sqlx.DB instance
func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
