package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	Id            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	AccountNumber string    `json:"account_number"`
	AccountName   string    `json:"account_name"`
	BankName      string    `json:"bank_name"`
	Balance       int       `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (account *Account) BeforeCreate(tx *gorm.DB) (err error) {
	account.Id = uuid.New()
	return
}
