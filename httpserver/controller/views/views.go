package views

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

const (
	M_BAD_REQUEST                      = "BAD_REQUEST"
	M_INVALID_CREDENTIALS              = "INVALID_CREDENTIALS"
	M_CREATED                          = "CREATED"
	M_ACCOUNT_VALIDATED                = "ACCOUNT VALID"
	M_OK                               = "OK"
	M_USERNAME_ALREADY_USED            = "USER_ALREADY_USED"
	M_INTERNAL_SERVER_ERROR            = "INTERNAL_SERVER_ERROR"
	M_TRANSACTION_SUCCESSFULLY_DELETED = "TRANSACTION_SUCCESSFULLY_DELETED"
	M_TRANSACTION_NOT_FOUND            = "TRANSACTION_NOT_FOUND"
	M_ACCOUNT_NOT_FOUND                = "ACCOUNT_NOT_FOUND"
	M_ACCOUNT_VALIDATION_FAILED        = "ACCOUNT VALIDATION FAILED"
	M_CALLBACK_NOT_FOUND               = "CALLBACK NOT FOUND"
)

func SuccessResponse(status int, message string, payload interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Payload: payload,
	}
}

func ErrorResponse(status int, message string, error error) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Error:   error.Error(),
	}
}

func WriteJsonResponse(ctx *gin.Context, res *Response) {
	ctx.JSON(res.Status, res)
}
