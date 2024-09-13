package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/julioc98/winterfell/internal/domain"
)

// UseCase represents an use case.
type UseCase interface {
	Webhook(req domain.WebhookRequest) error
}

// RestHandler represents a REST handler.
type RestHandler struct {
	r  *chi.Mux
	uc UseCase
}

// NewRestHandler creates a new RestHandler.
func NewRestHandler(r *chi.Mux, uc UseCase) *RestHandler {
	return &RestHandler{
		r:  r,
		uc: uc,
	}
}

// RegisterHandlers registers the handlers of the REST API.
func (h *RestHandler) RegisterHandlers() {
	// easter egg
	h.r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Winter is coming!"))
	})
	h.r.Post("/webhook", h.Webhook)
}
