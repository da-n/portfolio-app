# portfolio-app

portfolio-app is an account management simulation app. The back-end is built in Go using a Hexagonal Architecture. The front-end is in Vue.js. The purpose of the app is to generate order sheets for an account that is setup in a modelled portfolio. The front-end UI permits a withdrawal amount to be specified, validates this and displays an order sheet with sale instructions.

### Criteria

1. Must be able to see current value of modelled portfolio account.
1. Must be able to specify a withdrawal amount.
1. Must not be able to withdraw more than the current value of the portfolio.
1. Must create an order sheet with a sell instruction.

## Definitions and language

* Context: manage a portfolio account
* Language: modelled portfolio, asset, isin, portfolio, account, withdrawal, buy, invest, sell, raise, units, instruction, holding, fund, order sheet, investor, customer, value, balance
* Entities: Customer, Account, WithdrawalRequest, OrderSheet, Instruction, Portfolio, Asset
* Service: CustomerService, AccountService, PortfolioService, OrderSheetService
* Events: withdrawal request created, withdrawal request accepted, withdrawal request rejected, instruction created, order created,
* Repository: CustomerRepository, PortfolioRepository, AssetRepository, AccountRepository, OrderSheetRepository

## Resources and API

The following structure will be as follows:

* There can be many modelled portfolios
* A modelled portfolio can have many assets to make up 100% of the portfolio
* There can be many customers
* A customer can have many accounts
* An account is of type modelled portfolio
* An account has a value
* An account can have many withdrawal requests
* A valid withdrawal request will generate an order sheet
* An order sheet can have many instructions
* An account can have many order sheets
* Instructions can of type "BUY", "INVEST", "SELL", "RAISE"

## Installation

Aside from Go, Docker, the `docker-compose` command, and a compiler that can run `make` commands is required.

## Instructions

A Makefile is included to aid running and testing. To start the back-end (Go) application run:

```
make run
```

The back-end app is served from `http://localhost:8080`, see further down for an overview of the API endpoints and Postman collection instructions.

## Tests

To run tests, first generate mocks by running the following command:

```
make mock
```

Then tests can be run with:

```
make test
```

## Back-end API (Go)

The back-end app is an API written in Go. The following API endpoints are enabled:

| Method | Endpoint                                                            | Description                  | Name                    |
|--------|---------------------------------------------------------------------|------------------------------|-------------------------|
| GET    | /customers/{customer}                                               | Get a customer by ID         | GetCustomer             |
| GET    | /customers/{customer}/accounts                                      | List accounts for a customer | ListAccounts            |
| GET    | /customers/{customer}/accounts/{account}                            | Get an account by ID         | GetAccount              |
| POST   | /customers/{customer}/accounts/{account}/withdrawal-requests        | Create a withdrawal request  | CreateWithdrawalRequest |
| GET    | /portfolios/{portfolio}                                             | Get a portfolio by ID        | GetPortfolio            |
