package response

import (
	"errors"
	"net/http"
	"testing"

	mock_utils "github.com/riyanda432/belajar-authentication/mocks/infra/utils"
	"github.com/stretchr/testify/suite"
)

type IResponseClientSuite struct {
	suite.Suite
	response   IResponseClient
	httpWriter mock_utils.MockResponseWriter
}

func (s *IResponseClientSuite) SetupTest() {
	s.response = NewResponseClient()
	s.httpWriter = mock_utils.MockResponseWriter{}
}

func (s *IResponseClientSuite) TestHttpError_commonErr() {
	s.httpWriter.MockHeader(http.Header{})
	s.httpWriter.MockWriteHeader(500)
	s.httpWriter.MockWrite(1, nil)
	e := s.response.HttpError(
		&s.httpWriter,
		errors.New("mock error"),
	)
	s.Nil(e)
}

func (s *IResponseClientSuite) TestHttpError() {
	s.httpWriter.MockHeader(http.Header{})
	s.httpWriter.MockWriteHeader(500)
	s.httpWriter.MockWrite(1, nil)
	e := s.response.HttpError(&s.httpWriter, errors.New("mock error"))
	s.Nil(e)
}

func (s *IResponseClientSuite) TestJson() {
	s.httpWriter.MockHeader(http.Header{})
	s.httpWriter.MockWrite(1, nil)
	e := s.response.JSON(&s.httpWriter, "xxx", "xxx", s.response.BuildMeta(1, 10, 100))
	s.Nil(e)
}

func TestIResponseClient(t *testing.T) {
	suite.Run(t, &IResponseClientSuite{})
}
