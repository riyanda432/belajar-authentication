package errors

import (
	error_helper "github.com/riyanda432/belajar-authentication/src/helper"
)

func InitErrorDicts() *error_helper.ErrorDictionaries {
	return error_helper.RegisterErrorDictionaries(errorCodes, httpCodes)
}

