package mock_helper

import (
	"errors"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MockDate() time.Time {
	return time.Date(2022, 4, 4, 11, 0, 0, 0, time.UTC)
}

// func MockDatePointer() *time.Time {
// 	x := MockDate()
// 	return &x
// }

func MockStringPointer() *string {
	x := "xxx"
	return &x
}

func MockUrl() string {
	return "https://stackoverflow.com/"
}

func MockGormDB() (*gorm.DB, *sqlmock.Sqlmock, error) {
	var (
		gormDB *gorm.DB
	)
	db, mockDB, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	if db == nil {
		return nil, nil, errors.New("mock db is null")
	}

	if mockDB == nil {
		return nil, nil, errors.New("sqlmock is null")
	}
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	gormDB, err = gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, nil, err
	}

	if gormDB == nil {
		return nil, nil, errors.New("gorm db is null")
	}
	return gormDB, &mockDB, nil
}
