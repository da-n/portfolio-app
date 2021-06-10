package domain

import (
	"github.com/da-n/portfolio-app/errs"
	"github.com/golang/mock/gomock"
	"testing"
)

var mockAccountRepo *MockAccountRepository
var accountService AccountService

func setupAccountServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockAccountRepo = NewMockAccountRepository(ctrl)
	accountService = NewAccountService(mockAccountRepo)

	return func() {
		accountService = nil
		defer ctrl.Finish()
	}
}

func TestItShouldReturnAnErrorWhenItCannotListAccounts(t *testing.T) {
	teardown := setupAccountServiceTest(t)
	defer teardown()

	mockAccountRepo.EXPECT().FindAllAccounts(1).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	_, err := accountService.ListAccounts(1)

	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func TestItShouldReturnASliceOfAccounts(t *testing.T) {
	teardown := setupAccountServiceTest(t)
	defer teardown()

	accounts := []Account{
		{
			Id:           1,
			CustomerId:   1,
			CurrencyCode: "GBP",
			Balance:      20000000,
		},
		{
			Id:           2,
			CustomerId:   2,
			Balance:      20000000,
			CurrencyCode: "GBP",
		},
	}
	mockAccountRepo.EXPECT().FindAllAccounts(1).Return(accounts, nil)

	a, _ := accountService.ListAccounts(1)

	if a == nil {
		t.Error("expected Accounts to be returned, got nil")
	}
}

func TestItShouldReturnAnErrorWhenItCannotGetAnAccount(t *testing.T) {
	teardown := setupAccountServiceTest(t)
	defer teardown()

	mockAccountRepo.EXPECT().FindAccountById(0).Return(nil, errs.NewNotFoundError("Account could not be found"))

	_, err := accountService.GetAccount(0)

	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func TestItShouldReturnAnAccount(t *testing.T) {
	teardown := setupAccountServiceTest(t)
	defer teardown()

	account := Account{
		Id:           1,
		CustomerId:   1,
		Balance:      20000000,
		CurrencyCode: "GBP",
	}
	mockAccountRepo.EXPECT().FindAccountById(1).Return(&account, nil)

	a, _ := accountService.GetAccount(1)

	if a == nil {
		t.Error("expected Account to be returned, got nil")
	}
}
