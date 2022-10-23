package middleware

import (
	"context"
	"net/http"
	"testing"

	mock_utils "github.com/riyanda432/belajar-authentication/mocks/infra/utils"
	"github.com/stretchr/testify/suite"
)

type suiteMiddleWare struct {
	suite.Suite
	w mock_utils.MockResponseWriter
}

func (s *suiteMiddleWare) SetupTest() {
	s.w = mock_utils.MockResponseWriter{}
}

func (s *suiteMiddleWare) TestCheckMobileAppHeader_BuyerIdNil() {
	s.w.MockHeader(http.Header{})
	s.w.MockWriteHeader(500)
	s.w.MockWrite(1, nil)

	req, _ := http.NewRequestWithContext(context.Background(), "", "/", nil)
	c := CheckMobileAppHeader(&http.ServeMux{})
	c.ServeHTTP(&s.w, req)
}

func (s *suiteMiddleWare) TestCheckMobileAppHeader_UserIdNil() {
	s.w.MockHeader(http.Header{})
	s.w.MockWriteHeader(500)
	s.w.MockWrite(1, nil)

	req, _ := http.NewRequestWithContext(context.Background(), "", "/", nil)
	c := CheckMobileAppHeader(&http.ServeMux{})
	c.ServeHTTP(&s.w, req)
}

func (s *suiteMiddleWare) TestCheckMobileAppHeader_BuyerIdNotNumber() {
	s.w.MockHeader(http.Header{})
	s.w.MockWriteHeader(500)
	s.w.MockWrite(1, nil)

	req, _ := http.NewRequestWithContext(context.Background(), "", "/", nil)

	c := CheckMobileAppHeader(&http.ServeMux{})
	c.ServeHTTP(&s.w, req)
}

func (s *suiteMiddleWare) TestCheckMobileAppHeader_UserIdNotNumber() {
	s.w.MockHeader(http.Header{})
	s.w.MockWriteHeader(500)
	s.w.MockWrite(1, nil)

	req, _ := http.NewRequestWithContext(context.Background(), "", "/", nil)

	c := CheckMobileAppHeader(&http.ServeMux{})
	c.ServeHTTP(&s.w, req)
}

func (s *suiteMiddleWare) TestCheckMobileAppHeader() {
	s.w.MockHeader(http.Header{})
	s.w.MockWriteHeader(500)
	s.w.MockWrite(1, nil)

	req, _ := http.NewRequestWithContext(context.Background(), "", "/", nil)

	c := CheckMobileAppHeader(&http.ServeMux{})
	c.ServeHTTP(&s.w, req)
}
func (s *suiteMiddleWare) TestCheckModbileAppHeader() {
	s.w.MockHeader(http.Header{})
	s.w.MockWriteHeader(500)
	s.w.MockWrite(1, nil)

	req, _ := http.NewRequestWithContext(context.Background(), "", "/", nil)
	req.Header.Set(XPlaftorm, "mobile-apps")

	c := CheckMobileAppHeader(&http.ServeMux{})
	c.ServeHTTP(&s.w, req)
}

func TestMiddleware(t *testing.T) {
	suite.Run(t, &suiteMiddleWare{})
}
