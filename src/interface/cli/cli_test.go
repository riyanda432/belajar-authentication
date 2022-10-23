package main

import (
	"errors"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	olduTExclude := uTExclude
	defer func() {
		uTExclude = olduTExclude
		r := recover()
		assert.Nil(t, r)
	}()
	uTExclude = func() error {
		return nil
	}
	os.Args = []string{"cli.go", "unit-test-validation-with-excluded-folder"}
	main()
}

func Test_main_err(t *testing.T) {
	olduTExclude := uTExclude
	defer func() {
		uTExclude = olduTExclude
		r := recover()
		e, _ := r.(error)
		assert.EqualError(t, e, "ut error")
	}()
	uTExclude = func() error {
		return errors.New("ut error")
	}
	os.Args = []string{"cli.go", "unit-test-validation-with-excluded-folder"}
	main()
}
