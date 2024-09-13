// Package app implements the application layer.
package app

import (
	"fmt"
	"testing"

	"github.com/julioc98/winterfell/internal/domain"
)

type mockTransferer struct {
	err error
}

func NewTransferer(e error) *mockTransferer {
	return &mockTransferer{
		err: e,
	}
}

func (m *mockTransferer) Transfer(domain.Transfer) error {
	return m.err
}

func TestUseCase_Webhook(t *testing.T) {
	type fields struct {
		t Transferer
	}
	type args struct {
		req domain.WebhookRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid invoice credited event",
			fields: fields{
				t: NewTransferer(nil),
			},
			args: args{
				req: domain.WebhookRequest{
					Event: domain.Event{
						Subscription: "invoice",
						Log: domain.Log{
							Type: "credited",
							Invoice: domain.Invoice{
								Amount: 1000,
								ID:     "inv_123",
								Link:   "http://example.com/invoice/inv_123",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Non-invoice subscription event",
			fields: fields{
				t: NewTransferer(nil),
			},
			args: args{
				req: domain.WebhookRequest{
					Event: domain.Event{
						Subscription: "non-invoice",
						Log: domain.Log{
							Type: "credited",
							Invoice: domain.Invoice{
								Amount: 1000,
								ID:     "inv_123",
								Link:   "http://example.com/invoice/inv_123",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Non-credited event type",
			fields: fields{
				t: NewTransferer(nil),
			},
			args: args{
				req: domain.WebhookRequest{
					Event: domain.Event{
						Subscription: "invoice",
						Log: domain.Log{
							Type: "debited",
							Invoice: domain.Invoice{
								Amount: 1000,
								ID:     "inv_123",
								Link:   "http://example.com/invoice/inv_123",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Transfer error",
			fields: fields{
				t: NewTransferer(fmt.Errorf("transfer error")),
			},
			args: args{
				req: domain.WebhookRequest{
					Event: domain.Event{
						Subscription: "invoice",
						Log: domain.Log{
							Type: "credited",
							Invoice: domain.Invoice{
								Amount: 1000,
								ID:     "inv_123",
								Link:   "http://example.com/invoice/inv_123",
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewUseCase(tt.fields.t)
			if err := uc.Webhook(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Webhook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
