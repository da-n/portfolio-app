package domain

import (
	"github.com/da-n/portfolio-app/errs"
	"github.com/golang/mock/gomock"
	"testing"
)

var mockCustomerRepo *MockCustomerRepository
var customerService CustomerService

func setupCustomerServiceTest(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockCustomerRepo = NewMockCustomerRepository(ctrl)
	customerService = NewCustomerService(mockCustomerRepo)

	return func() {
		customerService = nil
		defer ctrl.Finish()
	}
}

func TestItShouldReturnAnErrorWhenItCannotGetACustomer(t *testing.T) {
	teardown := setupCustomerServiceTest(t)
	defer teardown()

	mockCustomerRepo.EXPECT().FindByCustomerId(0).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	_, err := customerService.GetCustomer(0)

	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func TestItShouldReturnACustomerResponseWhenItCanGetACustomer(t *testing.T) {
	teardown := setupCustomerServiceTest(t)
	defer teardown()

	customer := Customer{
		Id:        1,
		FirstName: "Theia",
		LastName:  "Parker",
		Email:     "theia@example.com",
		Password:  "password123",
	}
	mockCustomerRepo.EXPECT().FindByCustomerId(1).Return(&customer, nil)

	c, _ := customerService.GetCustomer(1)

	if c == nil {
		t.Error("expected Customer to be returned, got nil")
	}
}
