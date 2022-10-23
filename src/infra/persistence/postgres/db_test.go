package postgres

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mock_utils "github.com/riyanda432/belajar-authentication/mocks/infra/utils"
	"github.com/riyanda432/belajar-authentication/src/infra/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func mockDB() *sql.DB {
	db, _, _ := sqlmock.New()
	return db
}

func TestNew(t *testing.T) {
	oldGormOpen := gormOpen
	defer func() {
		gormOpen = oldGormOpen
	}()
	gormOpen = func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {
		return &gorm.DB{
			Config: &gorm.Config{
				ConnPool:  mockDB(),
				Dialector: dialector,
			},
		}, nil
	}

	n := New(config.SqlDbConf{}, logrus.New())
	assert.IsType(t, &PostgresDb{}, n)
}

func TestNew_DBPingErr(t *testing.T) {
	oldGormOpen := gormOpen
	defer func() {
		gormOpen = oldGormOpen
		if r := recover(); r != nil {
			err := r.(error)
			assert.EqualError(t, err, "sql: database is closed")
		}
	}()
	x := mockDB()
	gormOpen = func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {

		return &gorm.DB{
			Config: &gorm.Config{
				ConnPool:  x,
				Dialector: dialector,
			},
		}, nil
	}
	x.Close()
	n := New(config.SqlDbConf{}, logrus.New())
	assert.IsType(t, &PostgresDb{}, n)
}

func TestNew_DBErr(t *testing.T) {
	oldGormOpen := gormOpen
	defer func() {
		gormOpen = oldGormOpen
		recover()
	}()
	gormOpen = func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {
		return &gorm.DB{
			Config: &gorm.Config{
				ConnPool:  nil,
				Dialector: dialector,
			},
		}, nil
	}

	n := New(config.SqlDbConf{}, mock_utils.MockLogger())
	assert.IsType(t, &PostgresDb{}, n)
}

func TestNew_DBConfNotZero(t *testing.T) {
	oldGormOpen := gormOpen
	defer func() {
		gormOpen = oldGormOpen
	}()
	gormOpen = func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {
		return &gorm.DB{
			Config: &gorm.Config{
				ConnPool:  mockDB(),
				Dialector: dialector,
			},
		}, nil
	}

	n := New(config.SqlDbConf{
		MaxOpenConn:            5,
		MaxIdleConn:            5,
		MaxIdleTimeConnSeconds: 5,
		MaxLifeTimeConnSeconds: 5,
	}, logrus.New())
	assert.IsType(t, &PostgresDb{}, n)
}

func TestNew_IsProdTrue(t *testing.T) {
	oldGormOpen := gormOpen
	defer func() {
		gormOpen = oldGormOpen
	}()
	gormOpen = func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {
		return &gorm.DB{
			Config: &gorm.Config{
				ConnPool:  mockDB(),
				Dialector: dialector,
			},
		}, nil
	}
	n := New(config.SqlDbConf{}, logrus.New())
	assert.IsType(t, &PostgresDb{}, n)
}

func TestNew_GormOpenError(t *testing.T) {
	oldGormOpen := gormOpen
	defer func() {
		gormOpen = oldGormOpen
		if r := recover(); r != nil {
			err := r.(string)
			assert.Equal(t, err, "Failed to connect to database!")
		}
	}()
	gormOpen = func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {
		return nil, errors.New("gormOpen error")
	}
	New(config.SqlDbConf{}, logrus.New())

}
