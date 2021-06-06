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

func Test_it_should_return_an_error_when_it_cannot_list_accounts(t *testing.T) {
	// setup
	teardown := setupAccountServiceTest(t)
	defer teardown()

	// given
	mockAccountRepo.EXPECT().FindAllAccounts(int64(1)).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	// when
	_, err := accountService.ListAccounts(int64(1))

	// then
	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func Test_it_should_return_a_slice_of_accounts(t *testing.T) {
	// setup
	teardown := setupAccountServiceTest(t)
	defer teardown()

	// given
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

	// when
	a, _ := accountService.ListAccounts(int64(1))

	// then
	if a == nil {
		t.Error("expected Accounts to be returned, got nil")
	}
}

func Test_it_should_return_an_error_when_it_cannot_get_an_account(t *testing.T) {
	// setup
	teardown := setupAccountServiceTest(t)
	defer teardown()

	// given
	mockAccountRepo.EXPECT().FindAccountById(int64(0)).Return(nil, errs.NewNotFoundError("Account could not be found"))

	// when
	_, err := accountService.GetAccount(int64(0))

	// then
	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func Test_it_should_return_an_account(t *testing.T) {
	// setup
	teardown := setupAccountServiceTest(t)
	defer teardown()

	// given
	account := appdomain.Account{
		Id:          int64(1),
		CustomerId:  int64(1),
		AccountType: appdomain.AccountTypePortfolio,
		Balance:     int64(20000000),
	}
	mockAccountRepo.EXPECT().FindAccountById(int64(1)).Return(&account, nil)

	// when
	a, _ := accountService.GetAccount(int64(1))

	// then
	if a == nil {
		t.Error("expected Account to be returned, got nil")
	}
}
