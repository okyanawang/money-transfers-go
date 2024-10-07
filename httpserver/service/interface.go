package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/params" // Adjust this import path
	"github.com/okyanawang/money-transfer-go/httpserver/controller/views"  // Adjust this import path
)

type AccountSvc interface {
	ValidateAccount(ctx context.Context, request *params.ValidateAccountRequest) *views.Response
	CreateAccount(ctx context.Context, account *params.CreateAccountRequest) *views.Response
	GetAccountByNumber(ctx context.Context, accountNumber string) *views.Response
	// GetAccountById(ctx context.Context, id uuid.UUID) *views.Response
	// UpdateAccount(ctx context.Context, account *params.UpdateAccountRequest, id uuid.UUID) *views.Response
}

type TransactionSvc interface {
	Transfer(ctx context.Context, transferRequest *params.TransferRequest) *views.Response
	GetTransactionById(ctx context.Context, id uuid.UUID) *views.Response
	// GetTransactionsByAccountId(ctx context.Context, accountId uuid.UUID) *views.Response
	UpdateTransactionStatus(ctx context.Context, params *params.UpdateTransactionRequest, id uuid.UUID) *views.Response
}

type CallbackSvc interface {
	// ProcessCallback(ctx context.Context, callbackRequest *params.CallbackRequest) *views.Response
	HandleTransactionCallback(ctx context.Context, req *params.TransactionCallback) *views.Response
}
