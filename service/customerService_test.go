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

func Test_it_should_return_an_error_when_it_cannot_get_a_customer(t *testing.T) {
	// setup
	teardown := setupCustomerServiceTest(t)
	defer teardown()

	// given
	mockCustomerRepo.EXPECT().FindById(int64(0)).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	// when
	_, err := customerService.GetCustomer(int64(0))

	// then
	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func Test_it_should_return_a_customer_response_when_it_can_get_a_customer(t *testing.T) {
	// setup
	teardown := setupCustomerServiceTest(t)
	defer teardown()

	// given
	customer := appdomain.Customer{
		Id:        int64(1),
		FirstName: "Theia",
		LastName:  "Parker",
		Email:     "theia@example.com",
		Password:  "password123",
	}
	mockCustomerRepo.EXPECT().FindById(int64(1)).Return(&customer, nil)

	// when
	c, _ := customerService.GetCustomer(int64(1))

	// then
	if c == nil {
		t.Error("expected Customer to be returned, got nil")
	}
}
