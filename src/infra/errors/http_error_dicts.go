package errors

import (
	error_helper "github.com/riyanda432/belajar-authentication/src/helper"
	"net/http"
)

var httpCodes = map[error_helper.ErrorCode]int{
	DATA_INVALID:                         http.StatusBadRequest,
	INVALID_HEADER_X_PLATFORM_MOBILE_APP: http.StatusForbidden,
	INVALID_PAYLOAD_CREATE_USER:          http.StatusBadRequest,
	INVALID_USER:                         http.StatusBadRequest,
	USER_NOT_FOUND:                       http.StatusBadRequest,
	USER_WRONG_PASSWORD:                  http.StatusUnauthorized,
	FAILED_RETRIEVE_USER:                 http.StatusInternalServerError,
	FAILED_CREATE_USER:                   http.StatusInternalServerError,
	INVALID_REQUEST_RETRIEVE_USER:        http.StatusBadRequest,
	INVALID_REQUEST_CREATE_USER:          http.StatusBadRequest,
	INVALID_REQUEST_LOGIN_USER:           http.StatusBadRequest,
	INVALID_PAYLOAD_LOGIN_USER:           http.StatusBadRequest,
}
