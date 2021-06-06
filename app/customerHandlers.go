package app

import (
	"encoding/json"
	"github.com/da-n/portfolio-app/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (h CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId, err := strconv.ParseInt(vars["customer_id"], 10, 64)
	if err != nil {
		writeJsonResponse(w, http.StatusBadRequest, err.Error())
	}
	customer, appErr := h.service.GetCustomer(customerId)
	if appErr != nil {
		writeJsonResponse(w, appErr.Code, appErr.Message)
	} else {
		writeJsonResponse(w, http.StatusOK, customer)
	}
}

func writeJsonResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}
