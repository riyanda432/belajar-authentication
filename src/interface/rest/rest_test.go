package rest

import (
	"context"
	"errors"
	"os"
	"syscall"
	"testing"

	mock_rest "github.com/riyanda432/belajar-authentication/mocks/interface/rest"
	usecases "github.com/riyanda432/belajar-authentication/src/app"
	"github.com/riyanda432/belajar-authentication/src/infra/config"
	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	mockServer := &mock_rest.MockiHttpServer{}
	mockLogger := &mock_rest.MockiLogger{}
	mockServer.MockListenAndServe(nil)
	mockServer.MockSetKeepAlivesEnabled()
	mockServer.MockShutdown(nil)
	mockLogger.MockInfo()
	mockLogger.MockWarnf()
	mockLogger.MockPrintln()
	quit := make(chan os.Signal, 1)
	go func() {
		Start(context.Background(), &HttpServer{
			mockServer,
			mockLogger,
			":8080",
		}, quit)
	}()
	quit <- syscall.SIGQUIT
}

func TestStart_Error(t *testing.T) {
	mockServer := &mock_rest.MockiHttpServer{}
	mockLogger := &mock_rest.MockiLogger{}
	mockServer.MockListenAndServe(errors.New("serve error"))
	mockLogger.MockFatal(errors.New("serve error"))
	mockLogger.MockInfo()
	mockLogger.MockWarnf()
	mockServer.MockSetKeepAlivesEnabled()
	mockServer.MockShutdown(errors.New("shutdown error"))
	mockLogger.MockError(errors.New("shutdown error"))
	mockLogger.MockPrintln()
	quit := make(chan os.Signal, 1)
	go func() {
		Start(context.Background(), &HttpServer{
			mockServer,
			mockLogger,
			":8080",
		}, quit)
	}()
	quit <- syscall.SIGQUIT
}

func TestNew(t *testing.T) {
	c := New(config.HttpConf{
		XRequestID: "1",
		Timeout:    1,
	},
		false,
		&mock_rest.MockiLogger{},
		usecases.AllUseCases{},
		false,
	)
	assert.IsType(t, &HttpServer{}, c)
}

func TestNew_timeoutNil(t *testing.T) {
	mockLogger := &mock_rest.MockiLogger{}
	mockLogger.MockFatalf("invalid http timeout")
	New(config.HttpConf{
		XRequestID: "1",
	},
		false,
		mockLogger,
		usecases.AllUseCases{},
		false,
	)
}

// func TestNew_xRequestIDNil(t *testing.T) {
// 	mockLogger := &mock_rest.MockiLogger{}
// 	mockLogger.MockFatalf("invalid x-request-id")
// 	New(config.HttpConf{
// 		Timeout: 1,
// 	},
// 		false,
// 		mockLogger,
// 		usecases.AllUseCases{},
// 		false,
// 	)
// }
