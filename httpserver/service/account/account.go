package account

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

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

// ValidateAccount calls mock API to validate the account details
func (svc *accountSvc) ValidateAccount(ctx context.Context, account *params.ValidateAccountRequest) *views.Response {
	// Replace with your actual mock API URL
	mockAPIUrl := "https://mockapi.io/endpoint-to-validate-account"

	// Prepare the request to send to the mock API
	req, err := http.NewRequest("POST", mockAPIUrl, nil) // Add body if needed
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	defer resp.Body.Close()

	// Read and parse the response from the mock API
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	var validationResponse map[string]interface{}
	err = json.Unmarshal(body, &validationResponse)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// Example logic to handle validation status
	status := validationResponse["status"]
	if status != "valid" {
		return views.ErrorResponse(http.StatusBadRequest, views.M_ACCOUNT_VALIDATION_FAILED, errors.New("account validation failed"))
	}

	// Return success response if validation is successful
	return views.SuccessResponse(http.StatusOK, views.M_ACCOUNT_VALIDATED, validationResponse)
}

func NewAccountSvc(repo repository.AccountRepo) service.AccountSvc {
	return &accountSvc{
		repo: repo,
	}
}
