package httphandler

import (
	"net/http"

	"github.com/jakobsym/sol_service/token/internal/controller/token"
)

type Handler struct {
	ctrl *token.Controller
}

func New(ctrl *token.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// obtain token address(?id=mSoLzYCxHdYgdzU16g5QSh3i5K3z3KZK7ytfqcJm7So)
func (h *Handler) GetTokenDetails(w http.ResponseWriter, r *http.Request) {
	address := r.FormValue("address")
	if address == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// make call to Get() handle the errors

	// return the data via encode

}
