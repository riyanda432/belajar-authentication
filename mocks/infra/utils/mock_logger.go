package mock_utils

import (
	"io"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
)

type MockIoWriter struct {
	mock.Mock
}

// Write implements io.Writer
func (*MockIoWriter) Write(p []byte) (n int, err error) {
	return n, err
}

var _ io.Writer = &MockIoWriter{}

type MockFormatter struct {
	mock.Mock
}

// Format implements logrus.Formatter
func (*MockFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return nil, nil
}

var _ logrus.Formatter = &MockFormatter{}

func MockLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       &MockIoWriter{},
		Level:     logrus.FatalLevel,
		Formatter: &MockFormatter{},
		ExitFunc: func(e int) {
			panic("mock")
		},
	}
}
