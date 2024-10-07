package views

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id            uuid.UUID `json:"id"`
	AccountNumber string    `json:"account_number"`
	AccountName   string    `json:"account_name"`
	BankName      string    `json:"bank_name"`
	Balance       int       `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type CreateAccountResponse struct {
	Id uuid.UUID `json:"id"`
	Account
}

type GetAccountResponse struct {
	Account
}

type UpdateAccountResponse struct {
	Account
}

type ValidateAccountResponse struct {
	IsValid bool   `json:"is_valid"`
	Message string `json:"message,omitempty"`
}
