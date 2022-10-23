package data_transfer_object_mobile_app_v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserCreate_ValidateSuccess(t *testing.T) {
	vt := UserCreate{
		FullName:        "store",
		UserName:        "agus@gmail.com",
		Password:        "agus12345678910",
		ConfirmPassword: "agus12345678910",
	}
	dto := NewUserCreateDTO(vt.FullName, vt.UserName, vt.Password, vt.ConfirmPassword)
	err := dto.Validate()
	assert.NotNil(t, dto, err)
}

func TestUserCreate_ValidateUserName_NotValid(t *testing.T) {
	vt := UserCreate{
		FullName:        "store",
		UserName:        "agusgmailcom",
		Password:        "agus1235678910",
		ConfirmPassword: "agus12345678910",
	}
	dto := NewUserCreateDTO(vt.FullName, vt.UserName, vt.Password, vt.ConfirmPassword)
	err := dto.Validate()
	assert.NotNil(t, dto, err)
}

func TestUserLogin_ValidateSuccess(t *testing.T) {
	vt := UserLogin{
		UserName:        "agus@gmail.com",
		Password:        "agus12345678910",
	}
	dto := NewUserLoginDTO(vt.UserName, vt.Password)
	err := dto.Validate()
	assert.NotNil(t, dto, err)
}


func TestUserLogin_ValidateFailed(t *testing.T) {
	vt := UserLogin{
		
	}
	dto := NewUserLoginDTO(vt.UserName, vt.Password)
	err := dto.Validate()
	assert.NotNil(t, dto, err)
}
