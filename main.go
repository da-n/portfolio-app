package main

import (
	"fmt"
	"github.com/da-n/portfolio-app/app"
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/service"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

func main() {
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
	userRepositoryDb := domain.NewUserRepositoryDb(dbClient)

	// create handlers
	userHandlers := app.UserHandlers{service.NewUserService(userRepositoryDb)}
}

// envCheck verify all necessary environment variables are set to run application
func envCheck() error {
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
			return fmt.Errorf("Environment variable %s not defined. Terminating application...", k)
		}
	}

	return nil
}

// getDbClient create a new database client using sqlx
func getDbClient() (*sqlx.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client, nil
}
