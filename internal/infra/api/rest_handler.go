// Package api implements the API layer.
package api

import (
	"encoding/json"
	"net/http"

	"github.com/julioc98/winterfell/internal/domain"
)

// Webhook handles the webhook endpoint.
func (h *RestHandler) Webhook(w http.ResponseWriter, r *http.Request) {
	var request domain.WebhookRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err := h.uc.Webhook(request); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
