package postgres

import (
	infra_model "github.com/riyanda432/belajar-authentication/src/infra/models"
	"gorm.io/gorm"
)

var dbAutoMigrate = func(db *gorm.DB, dst ...interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if n, ok := r.(error); ok {
				err = n
			}
		}
	}()
	return db.AutoMigrate(dst...)
}

// Migrate represent migration schema models
func Migrate(db *gorm.DB) error {
	User := infra_model.User{}

	// auto migrate
	if err := dbAutoMigrate(
		db,
		&User,
	); err != nil {
		return err
	}

	return nil
}
