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
	ID 					   string  
	AccountID 		 string   
	Status         string   
	Amount         float64  
	Description    string   
  CardLastDigits string   
	CreatedAt      string   
	UpdatedAt      string   

}