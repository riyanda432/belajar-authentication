package postgres

import (
	"database/sql"
	"fmt"

	"github.com/riyanda432/belajar-authentication/src/infra/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
)

type PostgresDb struct {
	DB    *gorm.DB
	SqlDB *sql.DB
}

var gormOpen = gorm.Open

func New(conf config.SqlDbConf, logger *logrus.Logger) *PostgresDb {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host,
		conf.Username,
		conf.Password,
		conf.Name,
		conf.Port,
	)

	if conf.Password == "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=disable",
			conf.Host,
			conf.Username,
			conf.Name,
			conf.Port,
		)
	}

	db, err := gormOpen(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: log.Default.LogMode(log.Warn),
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("database err: %s", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	logger.Printf("sql database connection %s success", db.Name())

	return &PostgresDb{
		DB:    db,
		SqlDB: sqlDB,
	}
}
