package params

import (
	"time"

	"github.com/google/uuid"
)

type CallbackRequest struct {
	TransactionId uuid.UUID `json:"transaction_id" binding:"required"`
	Status        string    `json:"status" binding:"required"`
	Amount        float64   `json:"amount" binding:"required"`
	Currency      string    `json:"currency" binding:"required"`
	CallbackDate  time.Time `json:"callback_date"`
}

type TransactionCallback struct {
	TransactionId uuid.UUID `json:"transaction_id" validate:"required"`
	Status        string    `json:"status" validate:"required"`
	Amount        float64   `json:"amount"`
	ProcessedAt   time.Time `json:"processed_at,omitempty"`
	ReceivedAt    time.Time `json:"received_at,omitempty"`
}
