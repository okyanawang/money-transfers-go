package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/okyanawang/money-transfer-go/config"
	"github.com/okyanawang/money-transfer-go/httpserver"
	account_controller "github.com/okyanawang/money-transfer-go/httpserver/controller/account"
	callback_controller "github.com/okyanawang/money-transfer-go/httpserver/controller/callback"
	transaction_controller "github.com/okyanawang/money-transfer-go/httpserver/controller/transaction"
	"github.com/okyanawang/money-transfer-go/httpserver/repository/gorm"
	account_service "github.com/okyanawang/money-transfer-go/httpserver/service/account"
	callback_service "github.com/okyanawang/money-transfer-go/httpserver/service/callback"
	transaction_service "github.com/okyanawang/money-transfer-go/httpserver/service/transaction"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	router := gin.Default()

	accountRepo := gorm.NewAccountRepo(db)
	transactionRepo := gorm.NewTransactionRepo(db)
	callbackRepo := gorm.NewCallbackRepo(db)

	accountSvc := account_service.NewAccountSvc(accountRepo)
	transactionSvc := transaction_service.NewTransactionSvc(transactionRepo)
	callbackSvc := callback_service.NewCallbackSvc(callbackRepo)

	accountControl := account_controller.NewAccountController(accountSvc)
	transactionControl := transaction_controller.NewTransactionController(transactionSvc)
	callbackControl := callback_controller.NewCallbackController(callbackSvc)

	app := httpserver.NewRouter(router, *accountControl, *transactionControl, *callbackControl)

	app.Start(":8088")
}
