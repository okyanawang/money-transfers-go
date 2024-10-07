package params

import (
	"time"

	"github.com/google/uuid"
)

type TransferRequest struct {
	SenderAccountId   uuid.UUID `json:"sender_account_id" binding:"required"`
	ReceiverAccountId uuid.UUID `json:"receiver_account_id" binding:"required"`
	Amount            int       `json:"amount" binding:"required"`
	TransactionDate   time.Time `json:"transaction_date"`
}

type UpdateTransactionRequest struct {
	Id        uuid.UUID `json:"id"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetTransactionRequest struct {
	TransactionId uuid.UUID `json:"transaction_id" binding:"required"`
}
