package params

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	AccountNumber string    `json:"account_number" binding:"required"`
	AccountName   string    `json:"account_name" binding:"required"`
	Balance       int       `json:"balance" binding:"required"`
	BankName      string    `json:"bank_name" binding:"required"`
	CreatedAt     time.Time `json:"created_at"`
}

type ValidateAccountRequest struct {
	AccountNumber string `json:"account_number" binding:"required"`
	AccountName   string `json:"account_name" binding:"required"`
}

type UpdateAccountRequest struct {
	AccountName string `json:"account_name" binding:"required"`
	BankName    string `json:"bank_name" binding:"required"`
}

type GetAccountRequest struct {
	AccountId uuid.UUID `json:"account_id" binding:"required"`
}
