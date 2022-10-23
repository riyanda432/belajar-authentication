package usecases_mobile_app_v1

import (
	"testing"

	mock_repo_mobile_app "github.com/riyanda432/belajar-authentication/mocks/domain/repositories/mobile_app"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockIMobileAppContractsUseCase struct {
	mock.Mock
}

type MobileAppContractUseCaseSuite struct {
	suite.Suite
	usecase   IUserUseCase
	userRepo mock_repo_mobile_app.MockUserRepository
}

func (s *MobileAppContractUseCaseSuite) SetupTest() {
	s.userRepo = mock_repo_mobile_app.MockUserRepository{}
}

func (s *MobileAppContractUseCaseSuite) Test_MobileAppUseCase() {
	d := NewMobileAppUseCase(
		&s.userRepo,
	)
	s.NotNil(d)
}

func Test_IMobileAppUseCaseV1(t *testing.T) {
	suite.Run(t, new(MobileAppContractUseCaseSuite))
}
