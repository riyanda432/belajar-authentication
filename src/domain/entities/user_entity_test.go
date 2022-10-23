package domain_models

import (
	infra_constants "github.com/riyanda432/belajar-authentication/src/infra/constants"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func makeUserData() *User {
	d:= MakeUser(
		1, // id
		"agus supriyadi",
		"agus@gmail.com",
		"pass",
		time.Now(),
		time.Now(),
	)
	return d
}

func mockUserEntity() User {
	id := uint64(1)
	mockTime := time.Date(2020, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	return User{
		id:        id,
		fullname:  "Agus Supriyatno",
		username:  "agus@gmail.com",
		password:  "password",
		createdAt: mockTime,
		updatedAt: mockTime,
	}
}

func TestCreateUser(t *testing.T) {
	user := makeUserData()
	entity := CreateUser(
		user.fullname,
		user.username,
		user.password,
	)
	assert.NotNil(t, entity, "entity should be not nil")
}

func TestGetUserID(t *testing.T) {
	entity := mockUserEntity()
	result := entity.GetID()

	assert.Equal(t, uint64(1), result, "ID should be equal")
}

func TestGetPassword(t *testing.T) {
	entity := mockUserEntity()
	result := entity.GetPassword()

	assert.Equal(t, "password", result, "Password should be equal")
}

func TestGetFullName(t *testing.T) {
	entity := mockUserEntity()
	result := entity.GetFullname()

	assert.Equal(t, "Agus Supriyatno", result, "FullName should be equal")
}

func TestGetUserName(t *testing.T) {
	entity := mockUserEntity()
	result := entity.GetUsername()

	assert.Equal(t, "agus@gmail.com", result, "UserName should be equal")
}

func TestGetUserCreatedAt(t *testing.T) {
	entity := mockUserEntity()
	result := entity.GetCreatedAt()

	timeStr := result.Format(infra_constants.ISODateTimeFormat)
	assert.Equal(t, timeStr, "2020-01-01T00:00:00Z", "createdAt should be equal")
}

func TestGetUserUpdatedAt(t *testing.T) {
	entity := mockUserEntity()
	result := entity.GetUpdatedAt()

	timeStr := result.Format(infra_constants.ISODateTimeFormat)
	assert.Equal(t, timeStr, "2020-01-01T00:00:00Z", "createdAt should be equal")
}

func TestUser_SetFullName(t *testing.T) {
	e := mockUserEntity()

	e.SetFullName("Agus")
	assert.Equal(t, e.GetFullname(), "Agus", "it should be equal")
}

func TestUser_SetUserName(t *testing.T) {
	e := mockUserEntity()

	e.SetUserName("agus@gmail.com")
	assert.Equal(t, e.GetUsername(), "agus@gmail.com", "it should be equal")
}
