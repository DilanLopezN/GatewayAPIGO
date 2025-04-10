package entities

import "gateway/internals/domain"

const (
	StatusPending = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceInput struct {
	APIKey         string  
	AccountID      string  `json:"account_id"`
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentType    string  `json:"payment_type"`
	CardNumber     string  `json:"card_number"`
	CCV            string  `json:"ccv"`
	ExpiryMonth    int     `json:"expiry_month"`
	ExpiryYear     int     `json:"expiry_year"`
	CardHolderName string  `json:"cardholder_name"`
}

type InvoiceOutput struct {
	ID 						 string   `json:"id"`  
	AccountID 		 string   `json:"account_id"`
	Status         string   `json:"status"`
	Amount         float64  `json:"amount"`
	Description    string   `json:"description"`
	PaymentType    string   `json:"payment_type"`
  CardLastDigits string   `json:"card_last_digits"`
	CreatedAt      string   `json:"created_at"`
	UpdatedAt      string   `json:"updated_at"`
}

func ToInvoice(input *CreateInvoiceInput, accountId string) (*domain.Invoice, error) {
	creditCard := domain.CreditCard{
		Number:        input.CardNumber,
		CVV:           input.CCV,
		ExpiryMonth:   input.ExpiryMonth,
		ExpiryYear:    input.ExpiryYear,
		CardHolderName: input.CardHolderName,
	}

	invoice, err := domain.NewInvoice(accountId, input.Amount, input.Description, input.PaymentType, creditCard)
	if err != nil {
		return nil, err
	}

	return invoice, nil
}

func FromInvoice(invoice *domain.Invoice) InvoiceOutput {
	return InvoiceOutput{
		ID:             invoice.ID,
		AccountID:      invoice.AccountID,
		Status:         string(invoice.Status),
		Amount:         invoice.Amount,
		Description:    invoice.Description,
		PaymentType:    invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt:      invoice.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      invoice.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}	