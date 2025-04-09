package repository

import (
	"database/sql"
	"gateway/internals/domain"
	"time"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(account *domain.Account) error {

	stmt, err := r.db.Prepare(`INSERT INTO accounts (id, name, email, api_key, balance, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
}

defer stmt.Close()

	_, err = stmt.Exec(account.ID, account.NAME, account.Email, account.APIKey, account.Balance, account.CreatedAt, account.UpdatedAt)


	if err != nil {
		return err
	}

	return nil
}


func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {

	var account domain.Account
	var cratedAt, updatedAt time.Time


	err := r.db.QueryRow(`SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = $1`, apiKey).
		Scan(&account.ID, &account.NAME, &account.Email, &account.APIKey, &account.Balance, &cratedAt, &updatedAt)


		if err == sql.ErrNoRows {
			return nil, domain.ErrAccountNotFound
		}

		if err != nil {
			return nil, err
		}

		account.CreatedAt = cratedAt
		account.UpdatedAt = updatedAt	
		return &account, nil


}


func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {

	var account domain.Account
	var cratedAt, updatedAt time.Time


	err := r.db.QueryRow(`SELECT id, name, email, api_key, balance, created_at, updated_at FROM accounts WHERE api_key = $1`, id).
		Scan(&account.ID, &account.NAME, &account.Email, &account.APIKey, &account.Balance, &cratedAt, &updatedAt)


		if err == sql.ErrNoRows {
			return nil, domain.ErrAccountNotFound
		}

		if err != nil {
			return nil, err
		}

		account.CreatedAt = cratedAt
		account.UpdatedAt = updatedAt	
		return &account, nil


}

func (r *AccountRepository) UpdateBalance(account *domain.Account, amount float64) error {

	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	var currentBalance float64

	 err = tx.QueryRow(`SELECT balance FROM accounts WHERE id = $1 FOR UPDATE`, account.ID).Scan(&currentBalance)

	if err == sql.ErrNoRows {
		return domain.ErrAccountNotFound
	}

	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE accounts SET balance = $2, updated_at = $3 WHERE id = $1`, account.ID, currentBalance+amount, time.Now())

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *AccountRepository) Update(account *domain.Account) error {

	stmt, err := r.db.Prepare(`UPDATE accounts SET name = $1, email = $2, api_key = $3, balance = $4, updated_at = $5 WHERE id = $6`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(account.NAME, account.Email, account.APIKey, account.Balance, time.Now(), account.ID)


	if err != nil {
		return err
	}

	return nil
}