package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMigrate_Err(t *testing.T) {

	err := Migrate(&gorm.DB{})
	assert.EqualError(t, err, "runtime error: invalid memory address or nil pointer dereference")
}

func TestMigrate(t *testing.T) {
	olddbAutoMigrate := dbAutoMigrate
	defer func() {
		dbAutoMigrate = olddbAutoMigrate
	}()
	dbAutoMigrate = func(db *gorm.DB, dst ...interface{}) (err error) {
		return nil
	}
	err := Migrate(&gorm.DB{})
	assert.Nil(t, err)
}
