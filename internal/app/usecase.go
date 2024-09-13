// Package app implements the application layer.
package app

import (
	"fmt"

	"github.com/julioc98/winterfell/internal/domain"
)

// Effective Go: Interface names:
// By convention, one-method interfaces are named
// by the method name plus an -er suffix or similar modification to construct an agent noun:
// Reader, Writer, Formatter, CloseNotifier etc.

// Transferer realises the Transfer method.
type Transferer interface {
	Transfer(domain.Transfer) error
}

// UseCase represents a use case.
type UseCase struct {
	t Transferer
}

// NewUseCase creates a new UseCase.
func NewUseCase(t Transferer) *UseCase {
	return &UseCase{
		t: t,
	}
}

// Webhook is a use case for webhook.
func (uc *UseCase) Webhook(req domain.WebhookRequest) error {
	if req.Event.Subscription != "invoice" {
		return nil
	}

	switch req.Event.Log.Type {
	case "credited":
		return uc.transfer(domain.Transfer{
			Amount:      req.Event.Log.Invoice.Amount,
			ExternalID:  req.Event.Log.Invoice.ID,
			Description: fmt.Sprintf("Invoice Link: %s", req.Event.Log.Invoice.Link),
			// Use fix account data for now.
			Name:          "Stark Bank S.A.",
			TaxID:         "20.018.183/0001-80",
			BankCode:      "20018183", //Pix
			BranchCode:    "0001",
			AccountNumber: "6341320293482496",
			AccountType:   "payment",
		})
	default:
		return nil
	}
}

func (uc *UseCase) transfer(req domain.Transfer) error {
	err := uc.t.Transfer(req)
	if err != nil {
		return err
	}

	return nil
}
