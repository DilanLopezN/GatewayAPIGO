package services

import (
	"gateway/internals/domain"
	"gateway/internals/http/entities"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}

func (s *AccountService) CreateAccount(input entities.CreateAccountInput) (*entities.AccountOutput, error) {
	account := entities.ToAccount(input)


	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)

	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err

	}
	if existingAccount != nil {
		return nil, domain.ErrDuplicatedApiKey
	}

	err = s.repository.Save(account)

	if err != nil {
		return nil, err
	}

	 output := entities.FromAccount(account)
	return &output, nil

}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*entities.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)

	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := entities.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*entities.AccountOutput, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := entities.FromAccount(account)
	return &output, nil
}

func (s *AccountService)FindByID(id string) (*entities.AccountOutput, error) {
	account, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	output := entities.FromAccount(account)
	return &output, nil
}