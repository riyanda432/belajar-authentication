package handler_mobile_app_v1

import (
	// "bytes"
	"bytes"
	"context"
	"encoding/json"

	// "encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_usecases "github.com/riyanda432/belajar-authentication/mocks/app/v1/use_cases/mobile_app"
	mock_helper "github.com/riyanda432/belajar-authentication/mocks/helper"
	mock_response "github.com/riyanda432/belajar-authentication/mocks/interface/rest/response"
	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"

	request "github.com/riyanda432/belajar-authentication/src/interface/rest/v1/request/mobile_app"
	// "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"
)

type suiteIUserHandler struct {
	suite.Suite
	handler        IUserHandler
	responseClient mock_response.MockResponseIResponseClient
	usecase        mock_usecases.MockIUserUseCase
	w              http.ResponseWriter
	user           entities.User
}

func (s *suiteIUserHandler) SetupTest() {
	s.responseClient = mock_response.MockResponseIResponseClient{}
	s.usecase = mock_usecases.MockIUserUseCase{}
	s.handler = NewUserHandler(&s.responseClient, &s.usecase)
	s.w = httptest.NewRecorder()
	x:= entities.MakeUser(
		uint64(1),
		"Agus Supriyadi",
		"agus@gmail.com",
		"Test123456789",
		mock_helper.MockDate(),
		mock_helper.MockDate(),
	)
	s.user = *x
}

func (s *suiteIUserHandler) TestCreate_ErrNil() {
	jsonBytes, _ := json.Marshal(request.UserCreate{
		FullName: "Agus Supriyadi",
		UserName: "agus@gmail.com",
		Password: "agus123456789",
		ConfirmPassword: "agus123456789",
	})
	r, _ := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/",
		bytes.NewReader(jsonBytes),
	)
	r.Header.Set("x-device-id", "xxx")
	s.usecase.MockCreateSuccess(nil, nil)
	s.responseClient.MockHttpErrorSuccess()
	s.handler.Create(httptest.NewRecorder(), r)
}

func (s *suiteIUserHandler) TestLogin_ErrNil() {
	jsonBytes, _ := json.Marshal(request.UserLogin{
		UserName: "agus@gmail.com",
		Password: "agus123456789",
	})
	r, _ := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/",
		bytes.NewReader(jsonBytes),
	)
	r.Header.Set("x-device-id", "xxx")
	s.usecase.MockLoginSuccess(nil, nil)
	s.responseClient.MockHttpErrorSuccess()
	s.handler.Login(httptest.NewRecorder(), r)
}

func TestIUserHandler(t *testing.T) {
	suite.Run(t, &suiteIUserHandler{})
}
