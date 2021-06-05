package service

import (
	appdomain "github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/errs"
	"github.com/da-n/portfolio-app/mocks/domain"
	"github.com/golang/mock/gomock"
	"testing"
)

var mockRepo *domain.MockCustomerRepository
var service CustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockCustomerRepository(ctrl)
	service = NewCustomerService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_it_should_return_an_error_when_it_cannot_get_a_customer(t *testing.T) {
	// setup
	teardown := setup(t)
	defer teardown()

	// given
	customerId := "1"
	mockRepo.EXPECT().FindByCustomerId(customerId).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	// when
	_, err := service.GetCustomer(customerId)

	// then
	if err == nil {
		t.Error("expected err to be returned, got nil")
	}
}

func Test_it_should_return_a_customer_response_when_it_can_get_a_customer(t *testing.T) {
	// setup
	teardown := setup(t)
	defer teardown()

	// given
	customer := appdomain.Customer{
		CustomerId: "1",
		FirstName:  "Theia",
		LastName:   "Parker",
		Email:      "theia@example.com",
		Password:   "password123",
	}
	mockRepo.EXPECT().FindByCustomerId(customer.CustomerId).Return(&customer, nil)

	// when
	c, _ := service.GetCustomer(customer.CustomerId)

	// then
	if c == nil {
		t.Error("expected Customer to be returned, got nil")
	}
}
