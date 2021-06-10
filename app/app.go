package app

import (
	"encoding/json"
	"fmt"
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/errs"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	if err := envCheck(); err != nil {
		log.Fatal(err)
	}

	dbClient, err := getDbClient()
	if err != nil {
		log.Fatal(err)
	}

	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)

	accountHandlers := AccountHandlers{domain.NewAccountService(accountRepositoryDb)}
	customerHandlers := CustomerHandlers{domain.NewCustomerService(customerRepositoryDb)}

	router := mux.NewRouter()

	api := router.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.GetCustomer).Methods(http.MethodGet, http.MethodOptions).Name("GetCustomer")
	api.HandleFunc("/customers/{customer_id:[0-9]+}/accounts", accountHandlers.ListAccounts).Methods(http.MethodGet, http.MethodOptions).Name("ListAccounts")
	api.HandleFunc("/customers/{customer_id:[0-9]+}/accounts/{account_id:[0-9]+}", accountHandlers.GetAccount).Methods(http.MethodGet, http.MethodOptions).Name("GetAccount")
	api.HandleFunc("/customers/{customer_id:[0-9]+}/accounts/{account_id:[0-9]+}/withdrawal-requests", accountHandlers.CreateWithdrawalRequest).Methods(http.MethodPost, http.MethodOptions).Name("CreateWithdrawalRequest")
	api.HandleFunc("/portfolios/{portfolio_id:[0-9]+}", accountHandlers.GetPortfolio).Methods(http.MethodGet, http.MethodOptions).Name("GetPortfolio")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT")), setCorsHeaders(router)))
}

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

func setCorsHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// demo comments
		// for testing we are allowing all origins for CORS, normally this would be granular
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
