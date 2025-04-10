package domain

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusApproved     Status = "aprroved"
	StatusRejected   Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
	Status         Status
	Amount 					float64	
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	Number string
	CVV 	string
	ExpiryMonth int
	ExpiryYear int
	CardHolderName string
}


func NewInvoice(accountId string, amount float64, description string,
	paymentType string, creditCard CreditCard) (*Invoice, error) {
		
	if amount <= 0 {
		return nil, errors.New("invalid amount")
	}

	invoice := &Invoice{
		ID:             uuid.New().String(), // Replace with actual ID generation logic
		AccountID:      accountId,
		Status:         StatusPending,
		Amount:         amount,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: creditCard.Number[len(creditCard.Number)-4:], // Extract last 4 digits
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return invoice, nil
}


func (i *Invoice) Process() error {

	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	var newStatus Status

	if randomSource.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}


	i.Status = newStatus

	return nil
}


func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != StatusPending  {
		return ErrInvalidStatus
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()

	return nil
}
