package transformer_v1_mobile

import (
	"testing"
	"time"

	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
	// vo "github.com/riyanda432/belajar-authentication/src/domain/value_objects/order"
	infra_constant "github.com/riyanda432/belajar-authentication/src/infra/constants"
	"github.com/stretchr/testify/assert"
)

func mockUser() *entities.User {
	mockDate := time.Date(2022, 4, 4, 11, 0, 0, 0, time.Local)

	user := entities.MakeUser(
		uint64(1),
		"Agus Supriyadi",
		"agus@gmail.com",
		"agus1235678910",
		mockDate,
		mockDate,
	)

	return user
}

func TestTransformCreateUpdateUserNil(t *testing.T) {
	result := TransformCreateUpdate(nil)
	assert.Nil(t, result, "transform result should be nil")
}

func TestTransformCreateUpdate(t *testing.T) {
	vEntity := mockUser()
	r := TransformCreateUpdate(vEntity)
	expect := &UserTransformCreateUpdate{
		ID:        vEntity.GetID(),
		CreatedAt: vEntity.GetCreatedAt().Format(infra_constant.ISODateTimeFormat),
		UpdatedAt: vEntity.GetUpdatedAt().Format(infra_constant.ISODateTimeFormat),
	}
	assert.Equal(t, expect, r)
}

func TestTransformUserLoginNil(t *testing.T) {
	result := TransformSuccessLogin(nil)
	assert.Nil(t, result, "transform result should be nil")
}

func TestTransformUserLogin(t *testing.T) {
	vEntity := mockUser()
	r := TransformSuccessLogin(vEntity)
	expect := &UserTransformSuccessLogin{
		Username: vEntity.GetUsername(),
		LoginAt:  time.Now().Format(infra_constant.ISODateTimeFormat),
	}
	assert.Equal(t, expect, r)
}
