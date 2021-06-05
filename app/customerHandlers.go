package app

import (
	"encoding/json"
	"github.com/da-n/portfolio-app/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (h CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := h.service.GetCustomer(customerId)
	if err != nil {
		writeJsonResponse(w, err.Code, err.Message)
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