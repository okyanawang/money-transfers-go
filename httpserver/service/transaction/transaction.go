package transaction

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/params"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/views"
	"github.com/okyanawang/money-transfer-go/httpserver/repository"
	"github.com/okyanawang/money-transfer-go/httpserver/repository/models"
	"github.com/okyanawang/money-transfer-go/httpserver/service"
	"gorm.io/gorm"
)

type transactionSvc struct {
	repo repository.TransactionRepo
}

func (svc *transactionSvc) Transfer(ctx context.Context, params *params.TransferRequest) *views.Response {
	transaction := models.Transaction{
		SenderAccountId:   params.SenderAccountId,
		ReceiverAccountId: params.ReceiverAccountId,
		Amount:            params.Amount,
		Status:            "pending",
	}

	err := svc.repo.CreateTransaction(ctx, &transaction)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.Transaction{
		Id:                transaction.Id,
		SenderAccountId:   transaction.SenderAccountId,
		ReceiverAccountId: transaction.ReceiverAccountId,
		Amount:            transaction.Amount,
		Status:            transaction.Status,
		CreatedAt:         transaction.CreatedAt,
		UpdatedAt:         transaction.UpdatedAt,
	})
}

func (svc *transactionSvc) GetTransactionById(ctx context.Context, id uuid.UUID) *views.Response {
	transaction, err := svc.repo.GetTransactionById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_TRANSACTION_NOT_FOUND, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.Transaction{
		Id:                transaction.Id,
		SenderAccountId:   transaction.SenderAccountId,
		ReceiverAccountId: transaction.ReceiverAccountId,
		Amount:            transaction.Amount,
		Status:            transaction.Status,
		CreatedAt:         transaction.CreatedAt,
		UpdatedAt:         transaction.UpdatedAt,
	})
}

func (svc *transactionSvc) UpdateTransactionStatus(ctx context.Context, params *params.UpdateTransactionRequest, id uuid.UUID) *views.Response {
	transaction, err := svc.repo.GetTransactionById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	transaction.Status = params.Status

	err = svc.repo.UpdateTransactionStatus(ctx, id, transaction.Status)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.Transaction{
		Id:                transaction.Id,
		SenderAccountId:   transaction.SenderAccountId,
		ReceiverAccountId: transaction.ReceiverAccountId,
		Amount:            transaction.Amount,
		Status:            transaction.Status,
	})
}

func NewTransactionSvc(repo repository.TransactionRepo) service.TransactionSvc {
	return &transactionSvc{
		repo: repo,
	}
}
