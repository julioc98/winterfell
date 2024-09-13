// Package domain represents the domain layer of the application.
package domain

import "time"

type WebhookRequest struct {
	Event Event `json:"event,omitempty"`
}

type Descriptions struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Invoice struct {
	Amount             int            `json:"amount,omitempty"`
	Brcode             string         `json:"brcode,omitempty"`
	Created            time.Time      `json:"created,omitempty"`
	Descriptions       []Descriptions `json:"descriptions,omitempty"`
	DiscountAmount     int            `json:"discountAmount,omitempty"`
	Discounts          []any          `json:"discounts,omitempty"`
	DisplayDescription string         `json:"displayDescription,omitempty"`
	Due                time.Time      `json:"due,omitempty"`
	Expiration         int            `json:"expiration,omitempty"`
	Fee                int            `json:"fee,omitempty"`
	Fine               float64        `json:"fine,omitempty"`
	FineAmount         int            `json:"fineAmount,omitempty"`
	ID                 string         `json:"id,omitempty"`
	Interest           float64        `json:"interest,omitempty"`
	InterestAmount     int            `json:"interestAmount,omitempty"`
	Link               string         `json:"link,omitempty"`
	Name               string         `json:"name,omitempty"`
	NominalAmount      int            `json:"nominalAmount,omitempty"`
	Pdf                string         `json:"pdf,omitempty"`
	Rules              []any          `json:"rules,omitempty"`
	Splits             []any          `json:"splits,omitempty"`
	Status             string         `json:"status,omitempty"`
	Tags               []string       `json:"tags,omitempty"`
	TaxID              string         `json:"taxId,omitempty"`
	TransactionIds     []string       `json:"transactionIds,omitempty"`
	Updated            time.Time      `json:"updated,omitempty"`
}

type Log struct {
	Created time.Time `json:"created,omitempty"`
	Errors  []any     `json:"errors,omitempty"`
	ID      string    `json:"id,omitempty"`
	Invoice Invoice   `json:"invoice,omitempty"`
	Type    string    `json:"type,omitempty"`
}

type Event struct {
	Created      time.Time `json:"created,omitempty"`
	ID           string    `json:"id,omitempty"`
	Log          Log       `json:"log,omitempty"`
	Subscription string    `json:"subscription,omitempty"`
	WorkspaceID  string    `json:"workspaceId,omitempty"`
}
