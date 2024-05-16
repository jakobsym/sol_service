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

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {
	address := r.FormValue("address")
	if address == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
