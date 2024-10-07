package callback

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/okyanawang/money-transfer-go/httpserver/controller/params"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/views"
	"github.com/okyanawang/money-transfer-go/httpserver/repository"
	"github.com/okyanawang/money-transfer-go/httpserver/service"
)

type callbackSvc struct {
	repo repository.CallbackRepo
}

type transactionSvc struct {
	repo repository.TransactionRepo
}

func (svc *callbackSvc) HandleTransactionCallback(ctx context.Context, req *params.TransactionCallback) *views.Response {
	// Fetch the transaction based on the transaction ID
	transaction, err := svc.repo.GetCallbackByTransactionId(ctx, req.TransactionId)
	if err != nil {
		// If an error occurs, return an internal server error
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// Check if the transaction was found
	if transaction == nil {
		return views.ErrorResponse(http.StatusBadRequest, views.M_CALLBACK_NOT_FOUND, errors.New("transaction not found"))
	}

	// Process the callback by updating the transaction's status
	err = svc.repo.ProcessCallback(ctx, transaction.Id, req.Status)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.CallbackResponse{
		Id:            transaction.Id,
		TransactionId: transaction.Id,
		Status:        req.Status,
		ProcessedAt:   time.Now(),
	})
}

func NewCallbackSvc(repo repository.CallbackRepo) service.CallbackSvc {
	return &callbackSvc{
		repo: repo,
	}
}
