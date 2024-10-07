package views

import (
	"time"

	"github.com/google/uuid"
)

// Transaction represents the structure of a transaction response.
type Transaction struct {
	Id                uuid.UUID `json:"id"`
	SenderAccountId   uuid.UUID `json:"sender_account_id"`
	ReceiverAccountId uuid.UUID `json:"receiver_account_id"`
	Amount            int       `json:"amount"`
	Currency          string    `json:"currency"`
	Status            string    `json:"status"` // e.g., "pending", "completed", "failed"
	TransactionDate   time.Time `json:"transaction_date"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// CreateTransactionResponse represents the response after creating a transaction.
type CreateTransactionResponse struct {
	Id uuid.UUID `json:"id"`
	Transaction
}

// GetTransactionResponse represents the response for getting transaction details.
type GetTransactionResponse struct {
	Transaction
}

// UpdateTransactionResponse represents the response after updating a transaction.
type UpdateTransactionResponse struct {
	Transaction
}
