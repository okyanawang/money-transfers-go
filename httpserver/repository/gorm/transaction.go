package gorm

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/okyanawang/money-transfer-go/httpserver/repository"
	"github.com/okyanawang/money-transfer-go/httpserver/repository/models"
	"gorm.io/gorm"
)

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) repository.TransactionRepo {
	return &transactionRepo{db: db}
}

func (repo *transactionRepo) CreateTransaction(ctx context.Context, transaction *models.Transaction) error {
	transaction.Id = uuid.New()
	transaction.CreatedAt = time.Now()
	transaction.Status = "pending" // Default status for a new transaction
	return repo.db.WithContext(ctx).Create(transaction).Error
}

func (repo *transactionRepo) GetTransactionById(ctx context.Context, id uuid.UUID) (*models.Transaction, error) {
	transaction := new(models.Transaction)
	return transaction, repo.db.WithContext(ctx).Where("id = ?", id).Take(transaction).Error
}

func (repo *transactionRepo) UpdateTransactionStatus(ctx context.Context, id uuid.UUID, status string) error {
	return repo.db.WithContext(ctx).Model(&models.Transaction{}).Where("id = ?", id).Update("status", status).Error
}
