package app

import (
	"encoding/json"
	"fmt"
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/errs"
	"github.com/da-n/portfolio-app/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	// check that all env vars are set
	if err := envCheck(); err != nil {
		log.Fatal(err)
	}

	// get a mysql client instance
	dbClient, err := getDbClient()
	if err != nil {
		log.Fatal(err)
	}

	// create database repositories
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	portfolioRepositoryDb := domain.NewPortfolioRepositoryDb(dbClient)

	// create handlers
	accountHandlers := AccountHandlers{service.NewAccountService(accountRepositoryDb)}
	customerHandlers := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	portfolioHandlers := PortfolioHandlers{service.NewPortfolioService(portfolioRepositoryDb)}

	// routes
	router := mux.NewRouter()
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.GetCustomer).Methods(http.MethodGet).Name("GetCustomer")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/accounts", accountHandlers.ListAccounts).Methods(http.MethodGet).Name("ListAccounts")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/accounts/{account_id:[0-9]+}", accountHandlers.GetAccount).Methods(http.MethodGet).Name("GetAccount")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/accounts/{account_id:[0-9]+}/withdrawal-requests", accountHandlers.CreateWithdrawalRequest).Methods(http.MethodPost).Name("CreateWithdrawalRequest")
	router.HandleFunc("/portfolios/{portfolio_id:[0-9]+}", portfolioHandlers.GetPortfolio).Methods(http.MethodGet).Name("GetPortfolio")

	// start server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT")), router))
}

// envCheck verify all necessary environment variables are set to run application
func envCheck() *errs.AppError {
	envVars := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_ADDRESS",
		"DB_PORT",
		"DB_NAME",
	}

	for _, k := range envVars {
		if os.Getenv(k) == "" {
			return errs.NewUnexpectedError("Environment variable '" + k + "' not defined. Terminating application...")
		}
	}

	return nil
}

// getDbClient create a new database client using sqlx
func getDbClient() (*sqlx.DB, *errs.AppError) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client, nil
}

func writeJsonResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
