// package gateway implements the gateway layer.
package gateway

import (
	"fmt"

	"github.com/julioc98/winterfell/internal/domain"
	"github.com/starkbank/sdk-go/starkbank/transfer"
)

type TransferGateway struct {
}

func NewTransferGateway() *TransferGateway {
	return &TransferGateway{}
}

func (tg *TransferGateway) Transfer(req domain.Transfer) error {
	_, err := transfer.Create(
		[]transfer.Transfer{
			{
				Amount:        req.Amount,
				Name:          req.Name,
				TaxId:         req.TaxID,
				BankCode:      req.BankCode,
				BranchCode:    req.BranchCode,
				AccountNumber: req.AccountNumber,
				AccountType:   req.AccountType,
				ExternalId:    req.ExternalID,
				Description:   req.Description,
			},
		}, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			return fmt.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	return nil
}
