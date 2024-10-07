package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Callback struct {
	Id            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	TransactionId uuid.UUID `gorm:"type:uuid" json:"transaction_id"`
	Status        string    `json:"status"`
	Amount        float64   `json:"amount"`       // New field for transaction amount
	Currency      string    `json:"currency"`     // New field for currency
	ProcessedAt   time.Time `json:"processed_at"` // New field for the time the transaction was processed
	ReceivedAt    time.Time `json:"received_at"`
}

func (callback *Callback) BeforeCreate(tx *gorm.DB) (err error) {
	callback.Id = uuid.New()
	return
}
