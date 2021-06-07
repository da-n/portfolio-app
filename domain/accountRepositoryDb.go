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

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (r AccountRepositoryDb) FindAllAccounts(customerId int64) ([]Account, *errs.AppError) {
	statement := "select id, customer_id, portfolio_id, currency_code, balance from accounts where customer_id = ?"
	a := make([]Account, 0)
	err := r.client.Select(&a, statement, customerId)
	if err != nil {
		logger.Error("Error while querying accounts: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return a, nil
}

func (r AccountRepositoryDb) FindAccountById(accountId int64) (*Account, *errs.AppError) {
	statement := "select id, customer_id, portfolio_id, currency_code, balance from accounts where id = ?"
	var a Account
	err := r.client.Get(&a, statement, accountId)
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

func (r AccountRepositoryDb) SaveWithdrawalRequest(withdrawalRequest WithdrawalRequest) (*WithdrawalRequest, *errs.AppError) {
	createdAt := time.Now().Format(dbTSLayout)

	statement := "insert into withdrawal_requests (account_id, amount, created_at) values (?, ?, ?)"
	result, err := r.client.Exec(statement, withdrawalRequest.AccountId, withdrawalRequest.Amount, createdAt)
	if err != nil {
		logger.Error("Error while creating new withdrawal request: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new withdrawal request: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	withdrawalRequest.Id = id
	withdrawalRequest.CreatedAt = createdAt

	return &withdrawalRequest, nil
}

func (r AccountRepositoryDb) FindOrderSheetById(orderSheetId int64) (*OrderSheet, *errs.AppError) {
	statement := "select id, withdrawal_request_id, status, created_at from order_sheets where id = ?"
	var orderSheet OrderSheet
	err := r.client.Get(&orderSheet, statement, orderSheetId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Order sheet could not be found")
		} else {
			logger.Error("Error while querying order sheets: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &orderSheet, nil
}

func (r AccountRepositoryDb) SaveOrderSheet(orderSheet OrderSheet) (*OrderSheet, *errs.AppError) {
	createdAt := time.Now().Format(dbTSLayout)

	statement := "insert into order_sheets (withdrawal_request_id, status, created_at) values (?, ?, ?)"
	result, err := r.client.Exec(statement, orderSheet.WithdrawalRequestId, orderSheet.Status, createdAt)
	if err != nil {
		logger.Error("Error while creating new order sheet: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new order sheet: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	orderSheet.Id = id
	orderSheet.CreatedAt = createdAt

	return &orderSheet, nil
}

func (r AccountRepositoryDb) FindPortfolioById(portfolioId int64) (*Portfolio, *errs.AppError) {
	statement := "select id, name from portfolios where id = ?"
	var p Portfolio
	err := r.client.Get(&p, statement, portfolioId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Portfolio could not be found")
		} else {
			logger.Error("Error while querying portfolios: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	statement2 := "select a.id, a.isin, a.name, ap.percent from assets a join asset_portfolio ap on ap.asset_id = a.id and ap.portfolio_id = ?"
	a := make([]Asset, 0)
	assetErr := r.client.Select(&a, statement2, portfolioId)
	if assetErr != nil {
		logger.Error("Error while querying portfolio assets: " + assetErr.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	p.Assets = a

	return &p, nil
}

func (r AccountRepositoryDb) SaveInstruction(instruction Instruction) (*Instruction, *errs.AppError) {
	createdAt := time.Now().Format(dbTSLayout)

	statement := "insert into instructions (order_sheet_id, instruction_type, isin, amount, currency_code, created_at) values (?, ?, ?, ?, ?, ?)"
	result, err := r.client.Exec(statement, instruction.OrderSheetId, instruction.InstructionType, instruction.Isin, instruction.Amount, instruction.CurrencyCode, createdAt)
	if err != nil {
		logger.Error("Error while creating new instruction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new instruction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	instruction.Id = id
	instruction.CreatedAt = createdAt
	return &instruction, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
