package domain

import (
	"database/sql"
	"github.com/da-n/portfolio-app/errs"
	"github.com/da-n/portfolio-app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

// AccountRepositoryDb is the database implementation of AccountRepository
type AccountRepositoryDb struct {
	client *sqlx.DB
}

// FindAll finds all accounts belonging to a customer
func (r AccountRepositoryDb) FindAll(customerId int64) ([]Account, *errs.AppError) {
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
func (r AccountRepositoryDb) FindById(accountId int64) (*Account, *errs.AppError) {
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

// SaveWithdrawalRequest save a withdrawal request in the database
func (r AccountRepositoryDb) SaveWithdrawalRequest(req WithdrawalRequest) (*WithdrawalRequest, *errs.AppError) {
	sqlInsert := "insert into withdrawal_requests (account_id, amount, created_at) values (?, ?, ?)"

	result, err := r.client.Exec(sqlInsert, req.AccountId, req.Amount, time.Now().Format(dbTSLayout))
	if err != nil {
		logger.Error("Error while creating new withdrawal request: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new withdrawal request: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	req.Id = id
	return &req, nil
}

// NewAccountRepositoryDb instantiates a new AccountRepositoryDb passing in a sqlx.DB instance
func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
