package app

import (
	"encoding/json"
	"github.com/da-n/portfolio-app/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserHandlers struct {
	Service service.UserService
}

func (h UserHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	user, err := h.Service.GetUser(userId)
	if err != nil {
		writeJsonResponse(w, err.Code, err.Message)
	} else {
		writeJsonResponse(w, http.StatusOK, user)
	}
}

func writeJsonResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}