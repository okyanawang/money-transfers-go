package callback_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/params"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/views"
	"github.com/okyanawang/money-transfer-go/httpserver/service"
)

type CallbackController struct {
	svc service.CallbackSvc
}

func NewCallbackController(svc service.CallbackSvc) *CallbackController {
	return &CallbackController{
		svc: svc,
	}
}

func (control *CallbackController) HandleTransactionCallback(ctx *gin.Context) {
	var req params.TransactionCallback
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

	response := control.svc.HandleTransactionCallback(ctx, &req)
	views.WriteJsonResponse(ctx, response)
}
