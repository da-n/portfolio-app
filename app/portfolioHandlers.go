package app

import (
	"github.com/da-n/portfolio-app/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PortfolioHandlers struct {
	service service.PortfolioService
}

func (h PortfolioHandlers) GetPortfolio(w http.ResponseWriter, r *http.Request) {
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
