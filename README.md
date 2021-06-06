# portfolio-app

This is an account management app that simulates requesting withdrawals from a modelled portfolio account. The back-end is built in Go and follows a Hexagonal Architecture. The front-end is built in Vue.js. The purpose of the app is to generate order sheets for an account that is setup in a modelled portfolio. The front-end UI permits a withdrawal amount to be specified, validates it and displays an order sheet with sale instructions.

## Installation

Go is required to compile the app. `Docker` (with `docker-compose`) is required to run the database. `make` is not required but recommended, a `Makefile` is included that has several convenience helper methods for things like running the app and tests etc.

## Instructions

Build and run the app:

```
make run
```

Start the database in the separate tab/window:

```
cd resources/docker
docker-compose up -d
```

Navigate to [http://localhost:8080](http://localhost:8080) to start using the app.

## Tests

Note, mocks are not included in VCS, they will need to be generated in order to run tests. Generate mocks:

```
make mock
```

Run tests:

```
make test
```

## API

The following API resources are available:

| Method | Endpoint                                                         | Description                  | 
|--------|------------------------------------------------------------------|------------------------------|
| GET    | /api/customers/{customer}                                        | Get a customer by ID         |
| GET    | /api/customers/{customer}/accounts                               | List accounts for a customer |
| GET    | /api/customers/{customer}/accounts/{account}                     | Get an account by ID         |
| POST   | /api/customers/{customer}/accounts/{account}/withdrawal-requests | Create a withdrawal request  |
| GET    | /api/portfolios/{portfolio}                                      | Get a portfolio by ID        |

A Postman collection of the resources is available in `resources/postman`.

## Brief

1. Users must be able to see current value of their modelled portfolio account.
1. Users must be able to specify a withdrawal amount.
1. Users must not be able to withdraw more than the current value of the portfolio.
1. Users must receive an order sheet with sell instructions upon a successful withdrawal request.

## Definitions and language

* Context: manage a portfolio account
* Language: modelled portfolio, asset, isin, portfolio, account, withdrawal, buy, invest, sell, raise, units, instruction, holding, fund, order sheet, investor, customer, value, balance
* Entities: Customer, Account, WithdrawalRequest, OrderSheet, Instruction, Portfolio, Asset
* Service: CustomerService, AccountService, PortfolioService, OrderSheetService
* Events: withdrawal request created, withdrawal request accepted, withdrawal request rejected, instruction created, order created,
* Repository: CustomerRepository, PortfolioRepository, AssetRepository, AccountRepository, OrderSheetRepository

## Assumptions

The following assumptions are made:

* There can be many modelled portfolios
* A modelled portfolio can have many assets to make up 100% of the portfolio
* There can be many customers
* A customer can have many accounts
* An account is of type modelled portfolio
* An account has a balance
* An account can have many withdrawal requests
* A valid withdrawal request will generate an order sheet
* An account can have many order sheets
* An order sheet can have many instructions
* Instructions can of type "BUY", "INVEST", "SELL", "RAISE"

## Out of scope

* Authentication and authorization, the context is always customer 1
* Full CRUD API endpoints, only endpoints necessary for the brief are implemented
* Other types of instruction, only "SELL" is included
