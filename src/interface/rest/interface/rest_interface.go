package rest_interface

import "context"

type IHttpServer interface {
	ListenAndServe() error
	SetKeepAlivesEnabled(v bool)
	Shutdown(ctx context.Context) error
}

type ILogger interface {
	Fatal(args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Println(args ...interface{})
	Fatalf(format string, args ...interface{})
}
