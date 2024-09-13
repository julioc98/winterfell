// Package domain represents the domain layer of the application.
package domain

import "time"

type Transfer struct {
	ID                 string
	Amount             int
	Name               string
	TaxID              string
	BankCode           string
	BranchCode         string
	AccountNumber      string
	AccountType        string
	ExternalID         string
	Scheduled          *time.Time
	Description        string
	DisplayDescription string
	Tags               []string
	Fee                int
	Status             string
	TransactionIDs     []string
	Metadata           map[string]interface{}
	Created            *time.Time
	Updated            *time.Time
}
