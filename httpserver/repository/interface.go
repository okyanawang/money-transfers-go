package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/okyanawang/money-transfer-go/httpserver/repository/models"
)

type AccountRepo interface {
	CreateAccount(ctx context.Context, account *models.Account) error
	GetAccountByNumber(ctx context.Context, accountNumber string) (*models.Account, error)
}

type TransactionRepo interface {
	CreateTransaction(ctx context.Context, transaction *models.Transaction) error
	GetTransactionById(ctx context.Context, id uuid.UUID) (*models.Transaction, error)
	UpdateTransactionStatus(ctx context.Context, id uuid.UUID, status string) error
}

type CallbackRepo interface {
	ProcessCallback(ctx context.Context, id uuid.UUID, status string) error
	GetCallbackByTransactionId(ctx context.Context, transactionId uuid.UUID) (*models.Transaction, error)
	UpdateCallback(ctx context.Context, id uuid.UUID, status string) error
}
