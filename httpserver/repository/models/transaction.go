package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	Id                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	SenderAccountId   uuid.UUID `gorm:"type:uuid" json:"sender_account_id"`
	ReceiverAccountId uuid.UUID `gorm:"type:uuid" json:"receiver_account_id"`
	Amount            int       `json:"amount"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (transaction *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	transaction.Id = uuid.New()
	return
}
