package helper_error

import (
	"fmt"
	"net/http"
	"strings"

	ozzo_validation "github.com/go-ozzo/ozzo-validation"
)

type ErrorCode uint

type ValidationErrors map[string]string

type CommonError struct {
	ClientMessage    string           `json:"message"`
	SystemMessage    interface{}      `json:"data"`
	ValidationErrors ValidationErrors `json:"validationErrors,omitempty"`
	ErrorCode        ErrorCode        `json:"code"`
	ErrorMessage     *string          `json:"-"`
	ErrorTrace       *string          `json:"-"`
}

func (err CommonError) Error() string {
	return fmt.Sprintf("CommonError: %+v. Trace: %+v", err.ErrorMessage, err.ErrorTrace)
}

func buildValidationError(err error) ValidationErrors {
	var errors ValidationErrors = map[string]string{}

	errValidate := strings.Split(err.Error(), ";")
	for _, err := range errValidate {
		errPerField := strings.Split(err, ":")
		if len(errPerField[0]) <= 1 {
			errors["error"] = errPerField[0]
		} else {
			errors[strings.TrimSpace(errPerField[0])] = strings.TrimSpace(errPerField[1])
		}
	}

	return errors
}

func NewCommonError(errCode ErrorCode, err error) *CommonError {
	var errMsg *string
	var errTrace *string
	var clientMessage string = "Unknown error."
	var systemMessage interface{} = "Unknown error."
	var commonError = errorDicts.errorCodes[errCode]

	if err != nil {
		s := err.Error()
		errMsg = &s

		ss := fmt.Sprintf("%+v", err)
		errTrace = &ss

		if errCode == UNKNOWN_ERROR {
			systemMessage = ss
		}
	}

	if commonError == nil {
		return &CommonError{
			ClientMessage: clientMessage,
			SystemMessage: systemMessage,
			ErrorCode:     errCode,
			ErrorTrace:    errTrace,
			ErrorMessage:  errMsg,
		}
	}

	if _err, ok := err.(*CommonError); ok {
		return _err
	}

	return &CommonError{
		ClientMessage: commonError.ClientMessage,
		SystemMessage: commonError.SystemMessage,
		ErrorCode:     errCode,
		ErrorTrace:    errTrace,
		ErrorMessage:  errMsg,
	}
}

func (err *CommonError) SetClientMessage(message string) {
	err.ClientMessage = message
}

func (err *CommonError) SetSystemMessage(message interface{}) {
	err.SystemMessage = message
}

func (err *CommonError) SetValidationMessage(message interface{}) {
	if _err, ok := message.(ozzo_validation.Errors); ok {
		err.ValidationErrors = buildValidationError(_err)
	}
}

const UNKNOWN_ERROR ErrorCode = 0

var errorDicts *ErrorDictionaries = &ErrorDictionaries{
	errorCodes:  make(map[ErrorCode]*CommonError),
	httpCodes:   make(map[ErrorCode]int),
}

type ErrorDictionaries struct {
	errorCodes  map[ErrorCode]*CommonError
	httpCodes   map[ErrorCode]int
}

func RegisterErrorDictionaries(
	errorCodes map[ErrorCode]*CommonError,
	httpCodes map[ErrorCode]int,
) *ErrorDictionaries {
	errorDicts = &ErrorDictionaries{
		errorCodes:  errorCodes,
		httpCodes:   httpCodes,
	}

	errorDicts.errorCodes[UNKNOWN_ERROR] = &CommonError{
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	}

	errorDicts.httpCodes[UNKNOWN_ERROR] = http.StatusInternalServerError

	return errorDicts
}

type HttpError struct {
	CommonError
	HttpStatusNumber int    `json:"-"`
	HttpStatusName   string `json:"type"`
}

func (err HttpError) Error() string {
	return err.ClientMessage
}

func (err CommonError) GetHttpStatus() int {
	if errorDicts.httpCodes[err.ErrorCode] == 0 {
		return http.StatusInternalServerError
	}

	return errorDicts.httpCodes[err.ErrorCode]
}

func (err CommonError) ToHttpError() HttpError {
	httpStatusNumber := err.GetHttpStatus()

	return HttpError{
		CommonError:      err,
		HttpStatusNumber: httpStatusNumber,
		HttpStatusName:   GetHttpStatusText(httpStatusNumber),
	}
}

func GetHttpStatusText(httpStatus int) string {
	if text := http.StatusText(httpStatus); text != "" {
		upper := strings.ToUpper(text)
		return strings.ReplaceAll(upper, " ", "_")
	}

	return "INTERNAL_SERVER_ERROR"
}