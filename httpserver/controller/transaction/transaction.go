package transaction_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/params"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/views"
	"github.com/okyanawang/money-transfer-go/httpserver/service"
)

type TransactionController struct {
	svc service.TransactionSvc
}

func NewTransactionController(svc service.TransactionSvc) *TransactionController {
	return &TransactionController{
		svc: svc,
	}
}

func (control *TransactionController) Transfer(ctx *gin.Context) {
	var req params.TransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := validator.New().Struct(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := control.svc.Transfer(ctx, &req)
	views.WriteJsonResponse(ctx, response)
}

func (control *TransactionController) GetTransactionById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	transactionId, err := uuid.Parse(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid transaction ID format",
		})
		return
	}

	transactionResponse := control.svc.GetTransactionById(ctx, transactionId)
	if transactionResponse.Status != http.StatusOK {
		views.WriteJsonResponse(ctx, transactionResponse)
		return
	}

	views.WriteJsonResponse(ctx, transactionResponse)
}

func (control *TransactionController) UpdateTransaction(ctx *gin.Context) {
	idParam := ctx.Param("id")
	transactionId, err := uuid.Parse(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid transaction ID format",
		})
		return
	}

	var req params.UpdateTransactionRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	transactionResponse := control.svc.GetTransactionById(ctx, transactionId)
	if transactionResponse.Status != http.StatusOK {
		views.WriteJsonResponse(ctx, transactionResponse)
		return
	}

	response := control.svc.UpdateTransactionStatus(ctx, &req, transactionId)
	views.WriteJsonResponse(ctx, response)
}

/*
func (control *TransactionController) TransactionCallback(ctx *gin.Context) {
	var req params.CallbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate request
	err := validator.New().Struct(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := control.svc.TransactionCallback(ctx, &req)
	views.WriteJsonResponse(ctx, response)
}
*/
