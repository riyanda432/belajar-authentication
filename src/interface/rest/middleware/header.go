package middleware

import (
	"context"
	"errors"
	"net/http"

	error_helper "github.com/riyanda432/belajar-authentication/src/helper"
	commonError "github.com/riyanda432/belajar-authentication/src/infra/errors"
	"github.com/riyanda432/belajar-authentication/src/interface/rest/response"
)

const (
	XPlaftorm string = "x-platform"
)

const (
	MobileAppPlatform string = "mobile-app"
)

type ContextKey int

const (
	CtxHeader ContextKey = iota + 1
)

type ContexHeader struct {
	SellerId *uint64
	BuyerId  *uint64
	UserId   *uint64
}

func CheckMobileAppHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		xPlaftorm := r.Header.Get(XPlaftorm)

		if xPlaftorm == "" {
			err := errors.New("required x-platform")
			response.NewResponseClient().HttpError(w, err)
			return
		}

		if xPlaftorm != MobileAppPlatform {
			err := errors.New("mobile-app only access")
			cerr := error_helper.NewCommonError(commonError.INVALID_HEADER_X_PLATFORM_MOBILE_APP, err)
			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		ctx = context.WithValue(ctx, CtxHeader, nil)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
