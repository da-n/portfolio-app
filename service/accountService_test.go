package service

import (
	appdomain "github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/errs"
	"github.com/da-n/portfolio-app/mocks/domain"
	"github.com/golang/mock/gomock"
	"testing"
)

var mockAccountRepo *domain.MockAccountRepository
var accountService AccountService

func setupAccountServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockAccountRepo = domain.NewMockAccountRepository(ctrl)
	accountService = NewAccountService(mockAccountRepo)

	return func() {
		accountService = nil
		defer ctrl.Finish()
	}
}

func TestItShouldReturnAnErrorWhenItCannotListAccounts(t *testing.T) {
	teardown := setupAccountServiceTest(t)
	defer teardown()

	mockAccountRepo.EXPECT().FindAllAccounts(int64(1)).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	_, err := accountService.ListAccounts(int64(1))

	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func TestItShouldReturnASliceOfAccounts(t *testing.T) {
	teardown := setupAccountServiceTest(t)
	defer teardown()

	accounts := []appdomain.Account{
		{
			Id:          int64(1),
			CustomerId:  int64(1),
			AccountType: appdomain.AccountTypePortfolio,
			Balance:     int64(20000000),
		},
		{
			Id:          int64(2),
			CustomerId:  int64(2),
			AccountType: appdomain.AccountTypePortfolio,
			Balance:     int64(20000000),
		},
	}
	mockAccountRepo.EXPECT().FindAllAccounts(int64(1)).Return(accounts, nil)

	a, _ := accountService.ListAccounts(int64(1))

	if a == nil {
		t.Error("expected Accounts to be returned, got nil")
	}
}

func TestItShouldReturnAnErrorWhenItCannotGetAnAccount(t *testing.T) {
	teardown := setupAccountServiceTest(t)
	defer teardown()

	mockAccountRepo.EXPECT().FindAccountById(int64(0)).Return(nil, errs.NewNotFoundError("Account could not be found"))

	_, err := accountService.GetAccount(int64(0))

	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func TestItShouldReturnAnAccount(t *testing.T) {
	teardown := setupAccountServiceTest(t)
	defer teardown()

	account := appdomain.Account{
		Id:          int64(1),
		CustomerId:  int64(1),
		AccountType: appdomain.AccountTypePortfolio,
		Balance:     int64(20000000),
	}
	mockAccountRepo.EXPECT().FindAccountById(int64(1)).Return(&account, nil)

	a, _ := accountService.GetAccount(int64(1))

	if a == nil {
		t.Error("expected Account to be returned, got nil")
	}
}
