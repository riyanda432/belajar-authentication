package mock_utils

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockResponseWriter struct {
	mock.Mock
}

func (m *MockResponseWriter) MockHeader(res http.Header) {
	m.Mock.On("Header").Return(res)
}

// Header implements http.ResponseWriter
func (m *MockResponseWriter) Header() http.Header {
	args := m.Called()

	var res http.Header

	if n, ok := args.Get(0).(http.Header); ok {
		res = n
	}

	return res
}

func (m *MockResponseWriter) MockWrite(res int, err error) {
	m.Mock.On("Write", mock.Anything).Return(res, err)
}

// Write implements http.ResponseWriter
func (m *MockResponseWriter) Write(byte []byte) (int, error) {
	args := m.Called(byte)

	var res int
	var err error

	if n, ok := args.Get(0).(int); ok {
		res = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return res, err
}

func (m *MockResponseWriter) MockWriteHeader(statusCode int) {
	m.Mock.On("WriteHeader", statusCode)
}

// WriteHeader implements http.ResponseWriter
func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.Called(statusCode)
}

var _ http.ResponseWriter = &MockResponseWriter{}
