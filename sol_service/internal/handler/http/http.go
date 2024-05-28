package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jakobsym/sol_service/sol_service/internal/controller/sol_service"
)

type Handler struct {
	ctrl *sol_service.Controller
}

func New(ctrl *sol_service.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// handle GET /solservice requests
func (h *Handler) GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	address := r.FormValue("id")
	details, err := h.ctrl.Get(r.Context(), address)
	if err != nil && errors.Is(err, sol_service.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
