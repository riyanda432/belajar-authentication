package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	usecases "github.com/riyanda432/belajar-authentication/src/app"
	"github.com/riyanda432/belajar-authentication/src/infra/config"
	rest_interface "github.com/riyanda432/belajar-authentication/src/interface/rest/interface"
	route_v1 "github.com/riyanda432/belajar-authentication/src/interface/rest/v1/route"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)


// HttpServer holds the dependencies for a HTTP server.
type HttpServer struct {
	rest_interface.IHttpServer
	logger rest_interface.ILogger
	Addr   string
}

// New creates and configures a server serving all application routes.
// The server implements a graceful shutdown.
// chi.Mux is used for registering some convenient middlewares and easy configuration of
// routes using different http verbs.
func New(
	conf config.HttpConf,
	isProd bool,
	logger rest_interface.ILogger,
	useCases usecases.AllUseCases,
	debug bool,
) *HttpServer {
	// wrap all the routes
	r := chi.NewRouter()
	// apply all middleware before create any routes
	applyMiddleWare(r, conf.XRequestID, conf.Timeout, isProd, logger, debug)
	r = route_v1.MakeRouteV1(r, logger, useCases)

	// http service
	srv := http.Server{
		Addr:    ":" + conf.Port,
		Handler: r,
	}

	return &HttpServer{&srv, logger, ":" + conf.Port}
}

// RequestID ...
func RequestID(XRequestID string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			//override
			middleware.RequestIDHeader = XRequestID

			requestID := r.Header.Get(middleware.RequestIDHeader)
			if requestID == "" {
				requestID = uuid.New().String() //use google uuid
			}

			ctx = context.WithValue(ctx, middleware.RequestIDKey, requestID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func applyMiddleWare(
	r *chi.Mux,
	xRequestID string,
	timeout int,
	isProd bool,
	logger rest_interface.ILogger,
	debug bool,
) {
	// apply common middleware here ...

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// logging middleware
	if !isProd {
		r.Use(middleware.Logger)
	}

	// x-request-id middleware
	if len(xRequestID) <= 0 {
		logger.Fatalf("invalid x-request-id")
	} else {
		r.Use(RequestID(xRequestID))
	}

	// timeout middleware
	if timeout <= 0 {
		logger.Fatalf("invalid http timeout")
	} else {
		r.Use(RequestID(xRequestID))
	}

}

func Start(ctx context.Context, srv *HttpServer, quit chan os.Signal) {
	// run HTTP service
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.logger.Fatal(err)
		}
	}()

	// ready to serve
	srv.logger.Info("listen on", srv.Addr)

	srv.gracefulShutdown(ctx, quit)
}

func (srv *HttpServer) gracefulShutdown(ctx context.Context, quit chan os.Signal) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// wait forever until 1 of signals above are received
	<-quit
	srv.logger.Warnf("got signal: %v, shutting down server ...", quit)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		srv.logger.Error(err)
	}

	srv.logger.Println("server exiting")
}
