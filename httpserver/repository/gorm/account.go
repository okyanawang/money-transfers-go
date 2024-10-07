package gorm

import (
	"context"
	"time"

	"github.com/okyanawang/money-transfer-go/httpserver/repository"

	"github.com/okyanawang/money-transfer-go/httpserver/repository/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) repository.AccountRepo {
	return &accountRepo{db: db}
}

func (repo *accountRepo) CreateAccount(ctx context.Context, account *models.Account) error {
	account.Id = uuid.New()
	account.CreatedAt = time.Now()
	return repo.db.WithContext(ctx).Create(account).Error
}

func (repo *accountRepo) GetAccountByNumber(ctx context.Context, accountNumber string) (*models.Account, error) {
	account := new(models.Account)
	return account, repo.db.WithContext(ctx).Where("account_number = ?", accountNumber).Take(account).Error
}
