package app

import (
	"github.com/da-n/portfolio-app/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CustomerHandlers struct {
	service domain.CustomerService
}

func (h CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	customerId, err := strconv.ParseInt(vars["customer_id"], 10, 64)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, err.Error())
	}

	customer, appErr := h.service.GetCustomer(int(customerId))
	if appErr != nil {
		writeJsonResponse(w, appErr.Code, appErr.Message)
	} else {
		writeJsonResponse(w, http.StatusOK, customer)
	}
}
