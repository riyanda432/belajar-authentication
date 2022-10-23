package helper_error

import (
	"errors"
	"net/http"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/stretchr/testify/assert"
)

func TestOnlyNewClientErrorMessage(t *testing.T) {
	errMsg := NewCommonError(0, nil)
	errMsg.SetClientMessage("This is a new client message")

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, "This is a new client message", errMsg.ClientMessage, "Client message should be a new client message")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, errMsg.SystemMessage, "Unknown error.")
	}
}

func TestOnlyNewSystemErrorMessage(t *testing.T) {
	errMsg := NewCommonError(0, nil)
	errMsg.SetSystemMessage("This is a new system message")

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, errMsg.ClientMessage, "Unknown error.")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, "This is a new system message", errMsg.SystemMessage, "System message should be a new system message")
	}
}

func TestNewCommonErrorMessage(t *testing.T) {
	errMsg := NewCommonError(0, nil)
	errMsg.SetClientMessage("This is a new client message")
	errMsg.SetSystemMessage("This is a new system message")

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, "This is a new client message", errMsg.ClientMessage, "Client message should be a new client message")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, "This is a new system message", errMsg.SystemMessage, "System message should be a new system message")
	}
}

func TestCommonError(t *testing.T) {
	errMsg := NewCommonError(0, nil)

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, "Unknown error.", errMsg.ClientMessage, "Client message should be a new client message")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, "Unknown error.", errMsg.SystemMessage, "System message should be a new system message")
	}
}

func TestThrownError(t *testing.T) {
	errorDicts.errorCodes[0] = &CommonError{
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
	}

	errMsg := NewCommonError(0, errors.New("this is another error"))

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, "Unknown error.", errMsg.ClientMessage, "Client message should be a new client message")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, "Unknown error.", errMsg.SystemMessage, "System message should be a new system message")
	}

	if assert.NotNil(t, errMsg.ErrorMessage) {
		assert.Equal(t, "this is another error", *errMsg.ErrorMessage, "System message should be a new system message")
	}
}

func TestReThrowCommonError(t *testing.T) {
	errorDicts.errorCodes[1] = &CommonError{
		ClientMessage: "This is my client message",
		SystemMessage: "This is my system message",
		ErrorCode:     1,
	}
	errorDicts.errorCodes[2] = &CommonError{
		ClientMessage: "This is my second client message",
		SystemMessage: "This is my second system message",
		ErrorCode:     2,
	}

	errMsg := NewCommonError(1, errors.New("this is an error"))
	errMsg2 := NewCommonError(2, errMsg)

	if assert.NotNil(t, errMsg2.ClientMessage) {
		assert.Equal(t, "This is my client message", errMsg2.ClientMessage, "Client message should be \"This is my client message\"")
	}

	if assert.NotNil(t, errMsg2.SystemMessage) {
		assert.Equal(t, "This is my system message", errMsg2.SystemMessage, "System message should be \"This is my system message\"")
	}
}

func TestValidationErrors(t *testing.T) {
	type testStruct struct {
		name string
		age  int
	}

	value := &testStruct{name: "John", age: 20}
	errorValidation := validation.ValidateStruct(
		value,
		validation.Field(&value.name, validation.Required),
		validation.Field(&value.age, validation.Min(25)),
	)

	errMsg := NewCommonError(0, errorValidation)
	errMsg.SetValidationMessage(errorValidation)

	assert.NotNil(t, errMsg.ValidationErrors["age"])
}

func TestErrorString(t *testing.T) {
	commonError := NewCommonError(0, errors.New("this is internal error"))
	errMsg := commonError.Error()

	if assert.NotNil(t, errMsg) {
		assert.Contains(t, errMsg, "CommonError", "Trace")
	}
}

func TestGetKnownHttpStatus(t *testing.T) {
	errorDicts.httpCodes[1] = http.StatusBadRequest
	commonError := NewCommonError(1, nil)

	assert.Equal(t, http.StatusBadRequest, commonError.GetHttpStatus(), "It will return HTTP status 400")
}

func TestGetUnknownHttpStatus(t *testing.T) {
	delete(errorDicts.httpCodes, 1)
	commonError := NewCommonError(1, nil)

	assert.Equal(t, http.StatusInternalServerError, commonError.GetHttpStatus(), "It will return HTTP status 500")
}

func TestHttpErrorString(t *testing.T) {
	errorDicts.errorCodes[1] = &CommonError{
		ClientMessage: "this is a client message",
	}

	commonError := NewCommonError(1, nil)

	assert.Equal(t, "this is a client message", commonError.ToHttpError().Error(), "It will return string \"this is a client message\"")
}

func TestHttpErrorHasCommonError(t *testing.T) {
	commonError := NewCommonError(0, errors.New("this is an error"))

	assert.NotNil(t, commonError.ClientMessage, "HttpError should have ClientMessage")
	assert.NotNil(t, commonError.SystemMessage, "HttpError should have SystemMessage")
	assert.NotNil(t, commonError.ErrorCode, "HttpError should have ErrorCode")
	assert.NotNil(t, commonError.ErrorMessage, "HttpError should have ErrorMessage")
	assert.NotNil(t, commonError.ErrorTrace, "HttpError should have ErrorTrace")
}

