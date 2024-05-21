package httphandler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jakobsym/sol_service/token/internal/controller/token"
	model "github.com/jakobsym/sol_service/token/pkg"
)

type Handler struct {
	ctrl *token.Controller
}

func New(ctrl *token.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// obtain token address(?id=mSoLzYCxHdYgdzU16g5QSh3i5K3z3KZK7ytfqcJm7So)
func (h *Handler) GetTokenDetails(w http.ResponseWriter, r *http.Request) {
	address := r.FormValue("id")
	if address == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := h.ctrl.Get(r.Context(), model.TokenAddress(address))
	if err != nil && errors.Is(err, token.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
