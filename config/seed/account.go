package seed

import (
	"github.com/google/uuid"
	"github.com/okyanawang/money-transfer-go/httpserver/repository/models"
	"gorm.io/gorm"
)

// LoadAccounts seeds the database with initial account data
func LoadAccounts(db *gorm.DB) error {
	accounts := []models.Account{
		{
			Id:            uuid.New(),
			AccountNumber: "721539120",
			AccountName:   "Dena Natalia",
			BankName:      "Bank BCA",
			Balance:       100000,
		},
		{
			Id:            uuid.New(),
			AccountNumber: "631113302",
			AccountName:   "Cynthia Yaputera",
			BankName:      "Bank OCBC",
			Balance:       200000,
		},
	}

	// Seed accounts into the database
	for _, account := range accounts {
		if err := db.Create(&account).Error; err != nil {
			return err // Return error if any account already exists
		}
	}

	return nil
}
