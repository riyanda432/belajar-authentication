package domain_models

import (
	"time"
)

type User struct {
	id        uint64
	fullname  string
	username  string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func CreateUser(
	fullname string,
	username string,
	password string,
) *User {
	return &User{
		fullname:  fullname,
		username:  username,
		password:  password,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func MakeUser(
	id uint64,
	fullname string,
	username string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		id:        id,
		fullname:  fullname,
		username:  username,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (m *User) GetID() uint64 {
	return m.id
}

func (m *User) GetFullname() string {
	return m.fullname
}

func (m *User) GetUsername() string {
	return m.username
}

func (m *User) GetPassword() string {
	return m.password
}

func (m *User) GetCreatedAt() time.Time {
	return m.createdAt
}

func (m *User) GetUpdatedAt() time.Time {
	return m.updatedAt
}

func (m *User) SetUserName(userName string) {
	m.username = userName
}

func (m *User) SetFullName(fullName string) {
	m.fullname = fullName
}
