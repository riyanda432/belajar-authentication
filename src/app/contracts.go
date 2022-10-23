package usecases

import (
	mobile_app_v1_usecases "github.com/riyanda432/belajar-authentication/src/app/v1/use_cases/mobile_app"
	mobile_app_repositories "github.com/riyanda432/belajar-authentication/src/domain/repositories/mobile_app"
)

type AllUseCases struct {
	MobileAppV1UseCases mobile_app_v1_usecases.MobileAppV1UseCases
}

func NewAllUsecase(
	// mobile_app
	userRepository mobile_app_repositories.UserRepository,
) AllUseCases {
	return AllUseCases{
		MobileAppV1UseCases: mobile_app_v1_usecases.NewMobileAppUseCase(
			userRepository,
		),
	}
}
