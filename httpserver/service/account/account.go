package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/okyanawang/money-transfer-go/httpserver/controller/params"
	"github.com/okyanawang/money-transfer-go/httpserver/controller/views"
	"github.com/okyanawang/money-transfer-go/httpserver/repository"
	"github.com/okyanawang/money-transfer-go/httpserver/repository/models"
	"github.com/okyanawang/money-transfer-go/httpserver/service"
	"gorm.io/gorm"
)

type accountSvc struct {
	repo repository.AccountRepo
}

// CreateAccount implements service.AccountSvc.
func (svc *accountSvc) CreateAccount(ctx context.Context, account *params.CreateAccountRequest) *views.Response {
	param := models.Account{
		AccountNumber: account.AccountNumber,
		AccountName:   account.AccountName,
		BankName:      account.BankName,
		Balance:       account.Balance,
	}

	err := svc.repo.CreateAccount(ctx, &param)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.Account{
		Id:            param.Id,
		AccountNumber: param.AccountNumber,
		AccountName:   param.AccountName,
		BankName:      param.BankName,
		Balance:       param.Balance,
		CreatedAt:     param.CreatedAt,
		UpdatedAt:     param.UpdatedAt,
	})
}

func (svc *accountSvc) GetAccountByNumber(ctx context.Context, accountNumber string) *views.Response {
	account, err := svc.repo.GetAccountByNumber(ctx, accountNumber)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_ACCOUNT_NOT_FOUND, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.Account{
		Id:            account.Id,
		AccountNumber: account.AccountNumber,
		AccountName:   account.AccountName,
		BankName:      account.BankName,
		Balance:       account.Balance,
		CreatedAt:     account.CreatedAt,
		UpdatedAt:     account.UpdatedAt,
	})
}

func (svc *accountSvc) ValidateAccount(ctx context.Context, account *params.ValidateAccountRequest) *views.Response {
	mockAPIUrl := os.Getenv("MOCK_API_URL")
	if mockAPIUrl == "" {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, errors.New("mock API URL is not set"))
	}

	url := fmt.Sprintf("%s/accounts?account_number=%s", mockAPIUrl, account.AccountNumber)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// Parse the JSON response
	var accounts []map[string]interface{}
	err = json.Unmarshal(body, &accounts)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// Check if the account exists
	if len(accounts) == 0 {
		return views.ErrorResponse(http.StatusBadRequest, views.M_ACCOUNT_VALIDATION_FAILED, errors.New("account not found"))
	}

	accountInfo := accounts[0] // Get the first account info from the response

	// Return success with account details
	return views.SuccessResponse(http.StatusOK, views.M_ACCOUNT_VALIDATED, map[string]interface{}{
		"account_name": accountInfo["account_name"],
		"bank_name":    accountInfo["bank_name"],
	})
}

func NewAccountSvc(repo repository.AccountRepo) service.AccountSvc {
	return &accountSvc{
		repo: repo,
	}
}
