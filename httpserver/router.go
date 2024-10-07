package httpserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/okyanawang/money-transfer-go/common"
	account_controller "github.com/okyanawang/money-transfer-go/httpserver/controller/account"
	callback_controller "github.com/okyanawang/money-transfer-go/httpserver/controller/callback"
	transaction_controller "github.com/okyanawang/money-transfer-go/httpserver/controller/transaction"
)

type router struct {
	router      *gin.Engine
	account     account_controller.AccountController
	transaction transaction_controller.TransactionController
	callback    callback_controller.CallbackController
}

func NewRouter(r *gin.Engine, account account_controller.AccountController, transaction transaction_controller.TransactionController, callback callback_controller.CallbackController) *router {
	return &router{
		router:      r,
		account:     account,
		transaction: transaction,
		callback:    callback,
	}
}

func (r *router) Start(port string) {
	r.router.POST("/api/v1/accounts/validate", r.account.ValidateAccount)

	// r.router.POST("/api/v1/transfers", r.verifyToken, r.transaction.Transfer)
	r.router.POST("/api/v1/transfers", r.transaction.Transfer)

	r.router.POST("/api/v1/transfers-callback", r.callback.HandleTransactionCallback)

	r.router.Run(port)
}

func (r *router) verifyToken(ctx *gin.Context) {
	bearerToken := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")
	if len(bearerToken) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid bearer token",
		})
		return
	}
	claims, err := common.ValidateToken(bearerToken[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Set("userData", claims)
}
