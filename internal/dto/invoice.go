package dto

import (
	"time"

	"github.com/ViniAguiar1/go-gateway-api/internal/domain"
)

const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceInput struct {
	APIKey          string
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	PaymentType     string  `json:"payment_type"`
	CardNumber      string  `json:"card_number"`
	CardholderName  string  `json:"cardholder_name"`
	CardExpiryMonth int     `json:"expiry_month"`
	CardExpiryYear  int     `json:"expiry_year"`
	CardCVV         string  `json:"cvv"`
}

type InvoiceOutput struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoice(input CreateInvoiceInput, accountID string) (*domain.Invoice, error) {
	card := domain.CrediCard{
		Number:         input.CardNumber,
		CardholderName: input.CardholderName,
		ExpiryMonth:    input.CardExpiryMonth,
		ExpiryYear:     input.CardExpiryYear,
		CVV:            input.CardCVV,
	}
	return domain.NewInvoice(
		accountID,
		input.Amount,
		input.Description,
		input.PaymentType,
		card,
	)
}

func FromInvoice(invoice *domain.Invoice) InvoiceOutput {
	return InvoiceOutput{
		ID:             invoice.ID,
		AccountID:      invoice.AccountId,
		Amount:         invoice.Amount,
		Status:         string(invoice.Status),
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt,
		UpdatedAt:      invoice.UpdatedAt,
	}
}
