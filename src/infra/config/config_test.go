package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	oldosGetenv := osGetenv
	defer func() {
		osGetenv = oldosGetenv
	}()
	osGetenv = func(key string) string {
		if key == "HTTP_TIMEOUT" ||
			key == "DB_MAX_OPEN_CONN" ||
			key == "DB_MAX_IDLE_TIME_CONN_SECONDS" ||
			key == "DB_MAX_LIFE_TIME_CONN_SECONDS" ||
			key == "DB_MAX_IDLE_CONN" {
			return "1"
		} else {
			return oldosGetenv(key)
		}
	}
	c := Make()
	assert.IsType(t, Config{}, c)
}