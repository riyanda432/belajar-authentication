package usecases_mobile_app_v1

import (
	mobile_app_repositories "github.com/riyanda432/belajar-authentication/src/domain/repositories/mobile_app"
)

type MobileAppV1UseCases struct {
	UserUseCase IUserUseCase
}

func NewMobileAppUseCase(
	userRepo mobile_app_repositories.UserRepository,
) MobileAppV1UseCases {
	return MobileAppV1UseCases{
		UserUseCase: NewUserUseCase(
			userRepo,
		),
	}
}
