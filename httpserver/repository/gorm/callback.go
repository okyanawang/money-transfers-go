package gorm

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/okyanawang/money-transfer-go/httpserver/repository"
	"github.com/okyanawang/money-transfer-go/httpserver/repository/models"
	"gorm.io/gorm"
)

type callbackRepo struct {
	db *gorm.DB
}

func NewCallbackRepo(db *gorm.DB) repository.CallbackRepo {
	return &callbackRepo{db: db}
}

func (repo *callbackRepo) CreateCallback(ctx context.Context, callback *models.Callback) error {
	callback.Id = uuid.New()
	callback.ReceivedAt = time.Now()
	return repo.db.WithContext(ctx).Create(callback).Error
}

func (repo *callbackRepo) GetCallbackByTransactionId(ctx context.Context, transactionId uuid.UUID) (*models.Transaction, error) {
	transaction := new(models.Transaction)
	err := repo.db.WithContext(ctx).Where("id = ?", transactionId).Take(transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (repo *callbackRepo) UpdateCallback(ctx context.Context, id uuid.UUID, status string) error {
	return repo.db.WithContext(ctx).Model(&models.Callback{}).Where("id = ?", id).Update("status", status).Error
}

func (repo *callbackRepo) ProcessCallback(ctx context.Context, transactionId uuid.UUID, status string) error {
	return repo.db.WithContext(ctx).Model(&models.Transaction{}).
		Where("id = ?", transactionId).
		Updates(map[string]interface{}{
			"status": status,
		}).Error
}
