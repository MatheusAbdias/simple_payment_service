package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MatheusAbdias/simple_payment_service/domain/users"
)

type Controller struct {
	service *users.Service
}

func NewController(service *users.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	userDTO := users.NewUserDTO()
	if err := json.NewDecoder(r.Body).Decode(userDTO); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	user, err := controller.service.RegisterUser(r.Context(), userDTO)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	jsonResponse(w, user, http.StatusOK)
}

func handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Printf("Error: %s", err.Error())
	http.Error(w, err.Error(), statusCode)
}

func jsonResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		handleError(w, err, http.StatusInternalServerError)
	}
}
