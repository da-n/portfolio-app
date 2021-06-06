package domain

import (
	"database/sql"
	"github.com/da-n/portfolio-app/errs"
	"github.com/da-n/portfolio-app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// PortfolioRepositoryDb is the database implementation of PortfolioRepository
type PortfolioRepositoryDb struct {
	client *sqlx.DB
}

// FindPortfolioById find a portfolio by its ID
func (r PortfolioRepositoryDb) FindPortfolioById(portfolioId int64) (*Portfolio, *errs.AppError) {
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

// NewPortfolioRepositoryDb instantiates a new PortfolioRepositoryDb passing in a sqlx.DB instance
func NewPortfolioRepositoryDb(dbClient *sqlx.DB) PortfolioRepositoryDb {
	return PortfolioRepositoryDb{dbClient}
}
