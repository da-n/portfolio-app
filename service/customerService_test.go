package service

import (
	appdomain "github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/errs"
	"github.com/da-n/portfolio-app/mocks/domain"
	"github.com/golang/mock/gomock"
	"testing"
)

var mockCustomerRepo *domain.MockCustomerRepository
var customerService CustomerService

func setupCustomerServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockCustomerRepo = domain.NewMockCustomerRepository(ctrl)
	customerService = NewCustomerService(mockCustomerRepo)

	return func() {
		customerService = nil
		defer ctrl.Finish()
	}
}

func TestItShouldReturnAnErrorWhenItCannotGetACustomer(t *testing.T) {
	teardown := setupCustomerServiceTest(t)
	defer teardown()

	mockCustomerRepo.EXPECT().FindByCustomerId(int64(0)).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	_, err := customerService.GetCustomer(int64(0))

	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func TestItShouldReturnACustomerResponseWhenItCanGetACustomer(t *testing.T) {
	teardown := setupCustomerServiceTest(t)
	defer teardown()

	customer := appdomain.Customer{
		Id:        int64(1),
		FirstName: "Theia",
		LastName:  "Parker",
		Email:     "theia@example.com",
		Password:  "password123",
	}
	mockCustomerRepo.EXPECT().FindByCustomerId(int64(1)).Return(&customer, nil)

	c, _ := customerService.GetCustomer(int64(1))

	if c == nil {
		t.Error("expected Customer to be returned, got nil")
	}
}
