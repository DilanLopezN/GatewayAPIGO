package services

import (
	"gateway/internals/domain"
	"gateway/internals/http/entities"
)

type InvoiceService struct {
	InvoiceRepository domain.InvoiceRepository
	accountService AccountService
}

func NewInvoiceService(invoiceRepository domain.InvoiceRepository, accountService AccountService) *InvoiceService {
	return &InvoiceService{
		InvoiceRepository: invoiceRepository,
		accountService: accountService,
	}
}

func (s *InvoiceService) Create(input *entities.CreateInvoiceInput) (*entities.InvoiceOutput, error) {
	account, err := s.accountService.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := entities.ToInvoice(input, account.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountService.UpdateBalance(input.APIKey, invoice.Amount)
		if err != nil {
			return  nil, err
		}
	}


	err = s.InvoiceRepository.Save(invoice)
	if err != nil {
		return nil, err
	}


	output := entities.FromInvoice(invoice)
	return output, nil
}

func (s *InvoiceService) GetByID(id, apiKey string) (*entities.InvoiceOutput, error) {
	invoice, err := s.InvoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	account, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	if invoice.AccountID != account.ID {
		return nil, domain.ErrUnauthorized
	}


	output := entities.FromInvoice(invoice)
	return output, nil
}

func (s *InvoiceService) ListByAccount(accountID string) ([]*entities.InvoiceOutput, error) {
	invoices, err := s.InvoiceRepository.FindByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	output := make([]*entities.InvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		output[i] = entities.FromInvoice(invoice)
	}

	return output, nil
}

func (s *InvoiceService) ListByApiKey(apiKey string) ([]*entities.InvoiceOutput, error) {
	account, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccount(account.ID)
}
