package domain_repositories

import (
	"context"

	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
)

type UserRepository interface {
	Persist(context context.Context, user *entities.User) error
	DetailByUserName(
		ctx context.Context,
		UserName string,
	) (*entities.User, error)
}
