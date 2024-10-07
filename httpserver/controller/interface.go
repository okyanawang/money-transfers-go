package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type AccountController interface {
	CreateAccount(ctx *gin.Context)
	GetAccounts(ctx *gin.Context)
	GetAccountById(ctx *gin.Context)
	UpdateAccount(ctx *gin.Context)
	DeleteAccount(ctx *gin.Context)
	ValidateAccount(ctx *gin.Context)
}

type TransactionController interface {
	CreateTransaction(ctx *gin.Context)
	GetTransactions(ctx *gin.Context)
	GetTransactionById(ctx *gin.Context)
	UpdateTransaction(ctx *gin.Context)
	DeleteTransaction(ctx *gin.Context)
	TransactionCallback(ctx *gin.Context)
}

type CallbackController interface {
	HandleTransactionCallback(ctx *gin.Context)
}
