package views

import (
	"time"

	"github.com/google/uuid"
)

type CallbackResponse struct {
	Id            uuid.UUID `json:"id"`
	TransactionId uuid.UUID `json:"transaction_id"`
	Status        string    `json:"status"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	CallbackDate  time.Time `json:"callback_date"`
	ProcessedAt   time.Time `json:"processed_at"`
	ReceivedAt    time.Time `json:"received_at"`
}
