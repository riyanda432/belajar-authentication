package errors

import (
	error_helper "github.com/riyanda432/belajar-authentication/src/helper"
)

const (
	DATA_INVALID                         error_helper.ErrorCode = 78_099_00_00001
	INVALID_PAYLOAD_CREATE_USER          error_helper.ErrorCode = 78_099_00_00002
	INVALID_USER                         error_helper.ErrorCode = 78_099_00_00003
	USER_NOT_FOUND                       error_helper.ErrorCode = 78_099_00_00004
	FAILED_RETRIEVE_USER                 error_helper.ErrorCode = 78_099_00_00005
	FAILED_CREATE_USER                   error_helper.ErrorCode = 78_099_00_00006
	INVALID_REQUEST_RETRIEVE_USER        error_helper.ErrorCode = 78_099_00_00007
	INVALID_REQUEST_CREATE_USER          error_helper.ErrorCode = 78_099_00_00008
	INVALID_HEADER_X_PLATFORM_MOBILE_APP error_helper.ErrorCode = 78_099_00_00009
	INVALID_PAYLOAD_LOGIN_USER           error_helper.ErrorCode = 78_099_00_00010
	INVALID_REQUEST_LOGIN_USER           error_helper.ErrorCode = 78_099_00_00011
	USER_WRONG_PASSWORD                  error_helper.ErrorCode = 78_099_00_00012
)

var errorCodes = map[error_helper.ErrorCode]*error_helper.CommonError{
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},
	INVALID_PAYLOAD_CREATE_USER: {
		ClientMessage: "Failed to create user.",
		SystemMessage: "Request payload for create user has an invalid form.",
		ErrorCode:     INVALID_PAYLOAD_CREATE_USER,
	},
	INVALID_PAYLOAD_LOGIN_USER: {
		ClientMessage: "Failed to user login.",
		SystemMessage: "Request payload for user login has an invalid form.",
		ErrorCode:     INVALID_PAYLOAD_LOGIN_USER,
	},
	INVALID_USER: {
		ClientMessage: "Invalid user.",
		SystemMessage: "Username has been registered",
		ErrorCode:     INVALID_USER,
	},
	USER_WRONG_PASSWORD: {
		ClientMessage: "Invalid password",
		SystemMessage: "Password doesnt match",
		ErrorCode:     USER_WRONG_PASSWORD,
	},
	USER_NOT_FOUND: {
		ClientMessage: "Invalid user.",
		SystemMessage: "user not registered.",
		ErrorCode:     USER_NOT_FOUND,
	},
	FAILED_RETRIEVE_USER: {
		ClientMessage: "Failed to retrieve user.",
		SystemMessage: "Something wrong happened while retrieve user.",
		ErrorCode:     FAILED_RETRIEVE_USER,
	},
	FAILED_CREATE_USER: {
		ClientMessage: "Failed to create user.",
		SystemMessage: "Something wrong happened while create user.",
		ErrorCode:     FAILED_CREATE_USER,
	},
	INVALID_REQUEST_CREATE_USER: {
		ClientMessage: "Failed to create user.",
		SystemMessage: "Request has an invalid query params and/or payload to create user.",
		ErrorCode:     INVALID_REQUEST_CREATE_USER,
	},
	INVALID_REQUEST_LOGIN_USER: {
		ClientMessage: "Failed to user login.",
		SystemMessage: "Request has an invalid query params and/or payload to user login.",
		ErrorCode:     INVALID_REQUEST_LOGIN_USER,
	},
	INVALID_HEADER_X_PLATFORM_MOBILE_APP: {
		ClientMessage: "mobile-app only access",
		SystemMessage: "This endpoint can only be accessed from mobile-app platform",
		ErrorCode:     INVALID_HEADER_X_PLATFORM_MOBILE_APP,
	},
}
