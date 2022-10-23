package request_v1_mobile

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"net/http"
	"testing"

	error_helper "github.com/riyanda432/belajar-authentication/src/helper"
	common_error "github.com/riyanda432/belajar-authentication/src/infra/errors"

	"github.com/stretchr/testify/assert"
)

// user create test start
func TestUserCreate_ValidateSuccess(t *testing.T) {
	jsonBytes, _ := json.Marshal(UserCreate{
		FullName:        "Agus Supriyadi",
		UserName:        "agus@gmail.com",
		Password:        "agus12345678910",
		ConfirmPassword: "agus12345678910",
	})
	s := UserCreate{}
	r, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/",
		bytes.NewReader(jsonBytes),
	)
	if err != nil {
		t.Fatal(err)
	}
	res, e := s.Validate(r)
	assert.Nil(t, e)
	assert.NotNil(t, res)
}

func TestUserCreate_ValidateErr(t *testing.T) {
	jsonBytes, _ := json.Marshal(UserCreate{
		FullName:        "Agus Supriyadi",
		UserName:        "agusgmailcom",
		Password:        "agus12345678910",
		ConfirmPassword: "agus12345678910",
	})
	s := UserCreate{}
	r, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/",
		bytes.NewReader(jsonBytes),
	)
	if err != nil {
		t.Fatal(err)
	}
	res, e := s.Validate(r)
	assert.NotNil(t, e)
	assert.Nil(t, res)
}

func TestUserCreate_Validate_DecodeError(t *testing.T) {
	s := UserCreate{}
	r, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/",
		bytes.NewReader(nil),
	)
	if err != nil {
		t.Fatal(err)
	}
	res, e := s.Validate(r)
	assert.Equal(
		t,
		e,
		error_helper.NewCommonError(
			common_error.INVALID_PAYLOAD_CREATE_USER,
			errors.New("invalid JSON format"),
		),
	)
	assert.Nil(t, res)
}

// user create test end

// user login test start

func TestUserLogin_ValidateSuccess(t *testing.T) {
	jsonBytes, _ := json.Marshal(UserLogin{
		UserName: "agus@gmail.com",
		Password: "agus12345678910",
	})
	s := UserLogin{}
	r, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/",
		bytes.NewReader(jsonBytes),
	)
	if err != nil {
		t.Fatal(err)
	}
	res, e := s.Validate(r)
	assert.Nil(t, e)
	assert.NotNil(t, res)
}

func TestUserLogin_ValidateErr(t *testing.T) {
	jsonBytes, _ := json.Marshal(UserLogin{
		UserName: "",
		Password: "",
	})
	s := UserLogin{}
	r, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/",
		bytes.NewReader(jsonBytes),
	)
	if err != nil {
		t.Fatal(err)
	}
	res, e := s.Validate(r)
	assert.NotNil(t, e)
	assert.Nil(t, res)
}

func TestUserLogin_Validate_DecodeError(t *testing.T) {
	s := UserLogin{}
	r, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		"/",
		bytes.NewReader(nil),
	)
	if err != nil {
		t.Fatal(err)
	}
	res, e := s.Validate(r)
	assert.Equal(
		t,
		e,
		error_helper.NewCommonError(
			common_error.INVALID_PAYLOAD_LOGIN_USER,
			errors.New("invalid JSON format"),
		),
	)
	assert.Nil(t, res)
}

// user login test end
