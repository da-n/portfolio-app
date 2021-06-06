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

// FindAllAccounts finds all Account belonging to a customer
func (r AccountRepositoryDb) FindAllAccounts(customerId int64) ([]Account, *errs.AppError) {
	query := "select id, customer_id, portfolio_id, currency_code, balance from accounts where customer_id = ?"
	a := make([]Account, 0)
	err := r.client.Select(&a, query, customerId)
	if err != nil {
		logger.Error("Error while querying accounts: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return a, nil
}

// FindAccountById find an Account by its id
func (r AccountRepositoryDb) FindAccountById(accountId int64) (*Account, *errs.AppError) {
	query := "select id, customer_id, portfolio_id, currency_code, balance from accounts where id = ?"

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

// SaveWithdrawalRequest save a WithdrawalRequest in the database
func (r AccountRepositoryDb) SaveWithdrawalRequest(w WithdrawalRequest) (*WithdrawalRequest, *errs.AppError) {
	sqlInsert := "insert into withdrawal_requests (account_id, amount, created_at) values (?, ?, ?)"
	createdAt := time.Now().Format(dbTSLayout)

	result, err := r.client.Exec(sqlInsert, w.AccountId, w.Amount, createdAt)
	if err != nil {
		logger.Error("Error while creating new withdrawal request: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new withdrawal request: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	w.Id = id
	w.CreatedAt = createdAt
	return &w, nil
}

// FindOrderSheetById find an OrderSheet by its id
func (r AccountRepositoryDb) FindOrderSheetById(orderSheetId int64) (*OrderSheet, *errs.AppError) {
	query := "select id, withdrawal_request_id, status, created_at from order_sheets where id = ?"

	var o OrderSheet
	err := r.client.Get(&o, query, orderSheetId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Order sheet could not be found")
		} else {
			logger.Error("Error while querying order sheets: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &o, nil
}

// SaveWithdrawalRequest save a WithdrawalRequest in the database
func (r AccountRepositoryDb) SaveOrderSheet(o OrderSheet) (*OrderSheet, *errs.AppError) {
	sqlInsert := "insert into order_sheets (withdrawal_request_id, status, created_at) values (?, ?, ?)"
	createdAt := time.Now().Format(dbTSLayout)

	result, err := r.client.Exec(sqlInsert, o.WithdrawalRequestId, o.Status, createdAt)
	if err != nil {
		logger.Error("Error while creating new order sheet: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new order sheet: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	o.Id = id
	o.CreatedAt = createdAt
	return &o, nil
}

// FindPortfolioById find a Portfolio by its ID
func (r AccountRepositoryDb) FindPortfolioById(portfolioId int64) (*Portfolio, *errs.AppError) {
	query := "select id, name from portfolios where id = ?"
	var p Portfolio
	err := r.client.Get(&p, query, portfolioId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Portfolio could not be found")
		} else {
			logger.Error("Error while querying portfolios: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	// add the individual assets to the portfolio including the percentage of each one
	assetQuery := "select a.id, a.isin, a.name, ap.percent from assets a join asset_portfolio ap on ap.asset_id = a.id and ap.portfolio_id = ?"
	a := make([]Asset, 0)
	assetErr := r.client.Select(&a, assetQuery, portfolioId)
	if assetErr != nil {
		logger.Error("Error while querying portfolio assets: " + assetErr.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	p.Assets = a

	return &p, nil
}

// SaveInstruction save an Instruction in the database
func (r AccountRepositoryDb) SaveInstruction(i Instruction) (*Instruction, *errs.AppError) {
	sqlInsert := "insert into instructions (order_sheet_id, instruction_type, isin, amount, currency_code, created_at) values (?, ?, ?, ?, ?, ?)"
	createdAt := time.Now().Format(dbTSLayout)

	result, err := r.client.Exec(sqlInsert, i.OrderSheetId, i.InstructionType, i.Isin, i.Amount, i.CurrencyCode, createdAt)
	if err != nil {
		logger.Error("Error while creating new instruction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new instruction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	i.Id = id
	i.CreatedAt = createdAt
	return &i, nil
}

// NewAccountRepositoryDb instantiates a new AccountRepositoryDb passing in a sqlx.DB instance
func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
