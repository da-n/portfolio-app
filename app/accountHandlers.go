package app

import (
	"github.com/da-n/portfolio-app/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) ListAccounts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customerId"]
	accounts, err := h.service.ListAccounts(customerId)
	if err != nil {
		writeJsonResponse(w, err.Code, err.Message)
	} else {
		writeJsonResponse(w, http.StatusOK, accounts)
	}
}

func (h AccountHandlers) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customer, err := h.service.GetAccount(accountId)
	if err != nil {
		writeJsonResponse(w, err.Code, err.Message)
	} else {
		writeJsonResponse(w, http.StatusOK, customer)
	}
}
