package domain

import "errors"


var (
	ErrAccountNotFound = errors.New("account not found")

	ErrDuplicatedApiKey = errors.New("api Key already exists")

	ErrInvoiceNotFound = errors.New("invoice not found")

	ErrUnauthorized = errors.New("unauthorized access")
)