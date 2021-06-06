package app

import (
	"encoding/json"
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) ListAccounts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId, err := strconv.ParseInt(vars["customer_id"], 10, 64)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, err.Error())
	}
	accounts, appErr := h.service.ListAccounts(customerId)
	if appErr != nil {
		writeJsonResponse(w, appErr.Code, appErr.Message)
	} else {
		writeJsonResponse(w, http.StatusOK, accounts)
	}
}

func (h AccountHandlers) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId, err := strconv.ParseInt(vars["account_id"], 10, 64)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, err.Error())
	}
	customer, appErr := h.service.GetAccount(accountId)
	if appErr != nil {
		writeJsonResponse(w, appErr.Code, appErr.Message)
	} else {
		writeJsonResponse(w, http.StatusOK, customer)
	}
}

func (h AccountHandlers) CreateWithdrawalRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId, err := strconv.ParseInt(vars["account_id"], 10, 64)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, err.Error())
	}
	var req dto.WithdrawalRequestRequest
	if decodeErr := json.NewDecoder(r.Body).Decode(&req); decodeErr != nil {
		writeJsonResponse(w, http.StatusBadRequest, decodeErr.Error())
	} else {
		req.AccountId = accountId
		withdrawalRequest, appErr := h.service.CreateWithdrawalRequest(&req)
		if appErr != nil {
			writeJsonResponse(w, appErr.Code, appErr.Message)
		} else {
			writeJsonResponse(w, http.StatusOK, withdrawalRequest)
		}
	}
}

func (h AccountHandlers) GetOrderSheet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderSheetId, err := strconv.ParseInt(vars["order_sheet_id"], 10, 64)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, err.Error())
	}
	orderSheet, appErr := h.service.GetOrderSheet(orderSheetId)
	if appErr != nil {
		writeJsonResponse(w, appErr.Code, appErr.Message)
	} else {
		writeJsonResponse(w, http.StatusOK, orderSheet)
	}
}

func (h AccountHandlers) GetPortfolio(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portfolioId, err := strconv.ParseInt(vars["portfolio_id"], 10, 64)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, err.Error())
	}
	portfolio, appErr := h.service.GetPortfolio(portfolioId)
	if appErr != nil {
		writeJsonResponse(w, appErr.Code, appErr.Message)
	} else {
		writeJsonResponse(w, http.StatusOK, portfolio)
	}
}
