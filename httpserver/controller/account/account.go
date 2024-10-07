package account_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/params"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/views"
	"github.com/okyanawang/money-transfer-go/httpserver/service"
)

type AccountController struct {
	svc service.AccountSvc
}

func NewAccountController(svc service.AccountSvc) *AccountController {
	return &AccountController{
		svc: svc,
	}
}

func (control *AccountController) CreateAccount(ctx *gin.Context) {
	var req params.CreateAccountRequest
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

	response := control.svc.CreateAccount(ctx, &req)
	views.WriteJsonResponse(ctx, response)
}

/*
func (control *AccountController) GetAccountById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	accountId, err := uuid.Parse(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid account ID format",
		})
		return
	}

	accountResponse := control.svc.GetAccountById(ctx, accountId)
	if accountResponse.Status != http.StatusOK {
		views.WriteJsonResponse(ctx, accountResponse)
		return
	}

	views.WriteJsonResponse(ctx, accountResponse)
}

func (control *AccountController) UpdateAccount(ctx *gin.Context) {
	idParam := ctx.Param("id")
	accountId, err := uuid.Parse(idParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid account ID format",
		})
		return
	}

	var req params.UpdateAccountRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accountResponse := control.svc.GetAccountById(ctx, accountId)
	if accountResponse.Status != http.StatusOK {
		views.WriteJsonResponse(ctx, accountResponse)
		return
	}

	response := control.svc.UpdateAccount(ctx, &req, accountId)
	views.WriteJsonResponse(ctx, response)
}
*/

func (control *AccountController) ValidateAccount(ctx *gin.Context) {
	var req params.ValidateAccountRequest
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

	response := control.svc.ValidateAccount(ctx, &req)
	views.WriteJsonResponse(ctx, response)
}
