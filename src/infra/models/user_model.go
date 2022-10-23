package infra_model

import (
	"sync"
	"time"

	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
	"gorm.io/gorm/schema"
)

type User struct {
	ID        uint64    `json:"id" gorm:"column:id;primaryKey"`
	FullName  string    `json:"fullName" gorm:"column:full_name"`
	UserName  string    `json:"userName" gorm:"column:user_name"`
	Password  string    `json:"password" gorm:"column:password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (pm User) GetColumnName() []string {
	s, _ := schema.Parse(&User{}, &sync.Map{}, schema.NamingStrategy{})
	var m []string
	for _, field := range s.Fields {
		if field.DBName != "" {
			m = append(m, field.DBName)
		}
	}
	return m
}

func (u User) ToEntity() *entities.User {
	user := entities.MakeUser(
		u.ID,
		u.FullName,
		u.UserName,
		u.Password,
		u.CreatedAt,
		u.UpdatedAt,
	)
	return user
}

func ToUserModel(user entities.User) *User {
	return &User{
		ID:        user.GetID(),
		FullName:  user.GetFullname(),
		UserName:  user.GetUsername(),
		Password:  user.GetPassword(),
		CreatedAt: user.GetCreatedAt(),
		UpdatedAt: user.GetUpdatedAt(),
	}
}
