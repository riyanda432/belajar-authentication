package mock_handler

import (
	"net/http"

	handler "github.com/riyanda432/belajar-authentication/src/interface/rest/v1/handler/mobile_app"
)

type MockIUserHandler struct{}

// Create implements handler.IUserHandler
func (*MockIUserHandler) Create(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Login implements handler.IUserHandler
func (*MockIUserHandler) Login(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

var _ handler.IUserHandler = &MockIUserHandler{}
