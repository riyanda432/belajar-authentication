package main

import (
	"context"
	"database/sql"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
	infra_errors "github.com/riyanda432/belajar-authentication/src/infra/errors"

	"github.com/mattn/go-colorable"
	usecases "github.com/riyanda432/belajar-authentication/src/app"
	"github.com/riyanda432/belajar-authentication/src/infra/config"
	"github.com/riyanda432/belajar-authentication/src/infra/persistence/postgres"
	postgres_mobile_app "github.com/riyanda432/belajar-authentication/src/infra/persistence/postgres/mobile_app"
	"github.com/snowzach/rotatefilehook"

	"github.com/riyanda432/belajar-authentication/src/interface/rest"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

var postgresNew = postgres.New

var configMake = config.Make

func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := configMake()

	// initialize register error codes
	infra_errors.InitErrorDicts()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := NewLogInstance(
		LogName(conf.Log.Name),
		IsProduction(isProd),
		LogAdditionalFields(m))
	// open connection to persistence storage
	postgresdb := postgresNew(conf.SqlDb, logger)

	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.SqlDB, postgresdb.DB.Name())

	// Repositories and Usecases
	userRepository := postgres_mobile_app.NewUserRepository(postgresdb.DB)

	useCases := usecases.NewAllUsecase(
		userRepository,
	)

	// HTTP Handler
	// the server already implements a graceful shutdown.
	httpServer := rest.New(
		conf.Http,
		isProd,
		logger,
		useCases,
		conf.Debug,
	)

	quit := make(chan os.Signal, 1)
	rest.Start(ctx, httpServer, quit)
}
const Default = "default"

type DefaultFieldHook struct {
	fields map[string]interface{}
}

func (h *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultFieldHook) Fire(e *logrus.Entry) error {
	for i, v := range h.fields {
		e.Data[i] = v
	}
	return nil
}

type LogConfig struct {
	IsProduction bool
	LogFileName  string
	Fields       map[string]interface{}
}

type LogOption func(*LogConfig)

func IsProduction(isProd bool) LogOption {
	return func(o *LogConfig) {
		o.IsProduction = isProd
	}
}

func LogName(logname string) LogOption {
	return func(o *LogConfig) {
		o.LogFileName = logname
	}
}

func LogAdditionalFields(fields map[string]interface{}) LogOption {
	return func(o *LogConfig) {
		o.Fields = fields
	}
}

// NewLogInstance ...
func NewLogInstance(logOptions ...LogOption) *logrus.Logger {
	var level logrus.Level
	logger := logrus.New()

	//default configuration
	lc := &LogConfig{}
	lc.LogFileName = Default

	for _, opt := range logOptions {
		opt(lc)
	}

	//if it is production will output warn and error level
	if lc.IsProduction {
		level = logrus.WarnLevel
	} else {
		level = logrus.TraceLevel
	}

	logger.SetLevel(level)
	logger.SetOutput(colorable.NewColorableStdout())
	if lc.IsProduction {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
			PrettyPrint:     true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
			//PrettyPrint:     true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcname := s[len(s)-1]
				_, filename := path.Split(f.File)
				return funcname, filename
			},
		})
	}

	if !lc.IsProduction {
		dt := time.Now().UTC()
		rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
			Filename:   "logs/" + dt.Format("20060102") + "_" + lc.LogFileName,
			MaxSize:    50, // megabytes
			MaxBackups: 3,
			MaxAge:     28, //days
			Level:      level,
			Formatter: &logrus.JSONFormatter{
				TimestampFormat: time.RFC3339,
				CallerPrettyfier: func(f *runtime.Frame) (string, string) {
					s := strings.Split(f.Function, ".")
					funcname := s[len(s)-1]
					_, filename := path.Split(f.File)
					return funcname, filename
				},
			},
		})

		if err != nil {
			logger.Fatalf("Failed to initialize file rotate hook: %v", err)
		}

		logger.AddHook(rotateFileHook)
	}
	logger.AddHook(&DefaultFieldHook{lc.Fields})

	return logger
}
