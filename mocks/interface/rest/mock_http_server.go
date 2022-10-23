package mock_rest

import (
	"context"

	rest_interface "github.com/riyanda432/belajar-authentication/src/interface/rest/interface"
	"github.com/stretchr/testify/mock"
)

type MockiHttpServer struct {
	mock.Mock
}

func (m *MockiHttpServer) MockListenAndServe(e error) {
	m.Mock.On("ListenAndServe").Return(e)
}

// ListenAndServe implements iHttpServer
func (m *MockiHttpServer) ListenAndServe() error {
	args := m.Called()
	var e error
	if n, ok := args.Get(0).(error); ok {
		e = n
	}
	return e
}

func (m *MockiHttpServer) MockSetKeepAlivesEnabled() {
	m.Mock.On("SetKeepAlivesEnabled", mock.Anything)
}

// SetKeepAlivesEnabled implements iHttpServer
func (m *MockiHttpServer) SetKeepAlivesEnabled(v bool) {
	m.Called(v)
}

func (m *MockiHttpServer) MockShutdown(e error) {
	m.Mock.On("Shutdown", mock.Anything).Return(e)
}

// Shutdown implements iHttpServer
func (m *MockiHttpServer) Shutdown(ctx context.Context) error {
	args := m.Called(ctx)
	var e error
	if n, ok := args.Get(0).(error); ok {
		e = n
	}
	return e
}

var _ rest_interface.IHttpServer = &MockiHttpServer{}

type MockiLogger struct {
	mock.Mock
}

func (m *MockiLogger) MockFatalf(format string, args ...interface{}) {
	m.Mock.On("Fatalf", format, args)
}

// Fatalf implements iLogger
func (m *MockiLogger) Fatalf(format string, args ...interface{}) {
	m.Called(format, args)
}

func (m *MockiLogger) MockError(e error) {
	m.Mock.On("Error", e)
}

// Error implements iLogger
func (m *MockiLogger) Error(args ...interface{}) {
	m.Called(args...)
}

func (m *MockiLogger) MockFatal(e error) {
	m.Mock.On("Fatal", e)
}

// Fatal implements iLogger
func (m *MockiLogger) Fatal(args ...interface{}) {
	m.Called(args...)
}

func (m *MockiLogger) MockInfo() {
	m.Mock.On("Info", mock.Anything, mock.Anything)
}

// Info implements iLogger
func (m *MockiLogger) Info(args ...interface{}) {
	m.Called(args...)
}

func (m *MockiLogger) MockPrintln() {
	m.Mock.On("Println", mock.Anything)
}

// Println implements iLogger
func (m *MockiLogger) Println(args ...interface{}) {
	m.Called(args...)
}

func (m *MockiLogger) MockWarnf() {
	m.Mock.On("Warnf", mock.Anything, mock.Anything, mock.Anything)
}

// Warnf implements iLogger
func (m *MockiLogger) Warnf(format string, args ...interface{}) {
	m.Called(format, args)
}

var _ rest_interface.ILogger = &MockiLogger{}
