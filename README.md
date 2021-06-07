# portfolio-app

This is an account management app that simulates requesting withdrawals from a modelled portfolio account. The back-end is built in Go and follows a Hexagonal Architecture. The front-end is built in Vue.js. The purpose of the app is to generate order sheets for an account that is setup in a modelled portfolio. The front-end UI permits a withdrawal amount to be specified, validates it and displays an order sheet with sale instructions.

## Installation

There are a few dependencies:

* `Go` to compile and run the app
* `Docker` (with `docker-compose`) to run the database
* `Node` (with `NPM`) 
* `make` to more easily run the app and tests etc

## Instructions

There are three components to run the app, it is recommended to run these in three separate tabs/windows in a terminal: 

1. The back-end (`Go` binary)
1. The database (`Docker` with `docker-compose`)
1. The front-end (`NPM` run command)

To run the back-end:

```
make run
```
To run the database:

```
cd ./resources/docker
docker-compose up
```

To run the front-end for the first time:

```
cd ./front
npm ci
```

To run the front-end after installing node modules:

```
cd ./front
npm run serve
```

With everything up and running, navigate to the URL output from `npm run serve` (normally [http://localhost:8081](http://localhost:8081)).

## Tests

Mocks are not included in VCS, they need to be generated in order to run tests. 

Generate mocks:

```
make mock
```

With mocks generated can now run tests:

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

## App brief

1. Users must be able to see current value of their modelled portfolio account.
1. Users must be able to specify a withdrawal amount.
1. Users must not be able to withdraw more than the current value of the portfolio.
1. Users must receive an order sheet with sell instructions upon a successful withdrawal request.

## Definitions and language used in the app

* Context: manage a portfolio account
* Language: modelled portfolio, asset, isin, portfolio, account, withdrawal, buy, invest, sell, raise, units, instruction, holding, fund, order sheet, investor, customer, value, balance
* Entities: Customer, Account, WithdrawalRequest, OrderSheet, Instruction, Portfolio, Asset
* Service: CustomerService, AccountService, PortfolioService, OrderSheetService
* Events: withdrawal request created, withdrawal request accepted, withdrawal request rejected, instruction created, order created,
* Repository: CustomerRepository, PortfolioRepository, AssetRepository, AccountRepository, OrderSheetRepository

## Assumptions

The following assumptions exist:

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
