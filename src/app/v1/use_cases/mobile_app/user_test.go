package usecases_mobile_app_v1

import (
	"context"
	"errors"
	"testing"

	mock_dto_v1_mobile_app "github.com/riyanda432/belajar-authentication/mocks/app/v1/data_transfer_object/mobile_app"
	mock_domain "github.com/riyanda432/belajar-authentication/mocks/domain/repositories/mobile_app"
	mock_helper "github.com/riyanda432/belajar-authentication/mocks/helper"
	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
	"github.com/riyanda432/belajar-authentication/src/interface/rest/middleware"
	"github.com/stretchr/testify/suite"
)

type suiteIUserUseCase struct {
	suite.Suite
	usecase    IUserUseCase
	userRepo   mock_domain.MockUserRepository
	userEntity entities.User
	headers    middleware.ContexHeader
}

func (s *suiteIUserUseCase) SetupTest() {
	oe := entities.MakeUser(
		uint64(1),
		"Agus Supriyadi",
		"agus@gmail.com",
		"agus12345678912",
		mock_helper.MockDate(),
		mock_helper.MockDate(),
	)
	s.userEntity = *oe
	i := uint64(1)
	s.userRepo = mock_domain.MockUserRepository{}
	s.usecase = NewUserUseCase(&s.userRepo)
	s.headers.BuyerId = &i
}

// User Create Test Start
func (s *suiteIUserUseCase) Test_UserCreate_Success() {
	dtoUserCreate := &mock_dto_v1_mobile_app.MockUserDTO{}
	dtoUser := dtoUserCreate.MakeMockUserCreate()
	s.userRepo.MockPersist(&s.userEntity, nil)
	s.userRepo.MockDetailByUserNameSuccess(nil, nil)
	r, err := s.usecase.Create(context.Background(), dtoUser)
	s.NotNil(r)
	s.Nil(err)
}

func (s *suiteIUserUseCase) Test_UserCreate_PersistErr() {
	dtoUserCreate := &mock_dto_v1_mobile_app.MockUserDTO{}
	dtoUser := dtoUserCreate.MakeMockUserCreate()
	s.userRepo.MockDetailByUserNameSuccess(nil, nil)
	s.userRepo.MockPersistErr(nil, errors.New("d"))
	r, err := s.usecase.Create(context.Background(), dtoUser)
	s.Nil(r)
	s.NotNil(err)
}

func (s *suiteIUserUseCase) Test_UserCreate_UserHasBeenRegistered() {
	dtoUserCreate := &mock_dto_v1_mobile_app.MockUserDTO{}
	dtoUser := dtoUserCreate.MakeMockUserCreate()
	s.userRepo.MockPersist(&s.userEntity, nil)
	s.userRepo.MockDetailByUserNameSuccess(&s.userEntity, nil)
	r, err := s.usecase.Create(context.Background(), dtoUser)
	s.Nil(r)
	s.NotNil(err)
}

func (s *suiteIUserUseCase) Test_UserCreate_DTOErr() {
	s.userRepo.MockPersist(&s.userEntity, nil)
	s.userRepo.MockDetailByUserNameSuccess(&s.userEntity, nil)
	r, err := s.usecase.Create(context.Background(), nil)
	s.Nil(r)
	s.NotNil(err)
}

func (s *suiteIUserUseCase) Test_UserCreate_DetailByUserNameErr() {
	dtoUserCreate := &mock_dto_v1_mobile_app.MockUserDTO{}
	dtoUser := dtoUserCreate.MakeMockUserCreate()
	s.userRepo.MockPersistErr(nil, errors.New("d"))
	s.userRepo.MockDetailByUserNameErr(nil, errors.New("d"))
	r, err := s.usecase.Create(context.Background(), dtoUser)
	s.Nil(r)
	s.NotNil(err)
}

// User Create Test END

// User Login Test START
func (s *suiteIUserUseCase) Test_UserLogin_PassNotMatch() {
	dtoUserLogin := &mock_dto_v1_mobile_app.MockUserDTO{}
	dtoUser := dtoUserLogin.MakeMockUserLogin()
	s.userRepo.MockDetailByUserNameSuccess(&s.userEntity, nil)
	r, err := s.usecase.Login(context.Background(), dtoUser)
	s.Nil(r)
	s.NotNil(err)
}

func (s *suiteIUserUseCase) Test_UserLogin_DTO_Err() {
	s.userRepo.MockDetailByUserNameSuccess(&s.userEntity, nil)
	r, err := s.usecase.Login(context.Background(), nil)
	s.Nil(r)
	s.NotNil(err)
}

func (s *suiteIUserUseCase) Test_UserLogin_DetailByUserName_Err() {
	dtoUserLogin := &mock_dto_v1_mobile_app.MockUserDTO{}
	dtoUser := dtoUserLogin.MakeMockUserLogin()
	s.userRepo.MockDetailByUserNameErr(nil, errors.New("d"))
	r, err := s.usecase.Login(context.Background(), dtoUser)
	s.Nil(r)
	s.NotNil(err)
}

func (s *suiteIUserUseCase) Test_UserLogin_UserNotRegistered() {
	dtoUserLogin := &mock_dto_v1_mobile_app.MockUserDTO{}
	dtoUser := dtoUserLogin.MakeMockUserLogin()
	s.userRepo.MockDetailByUserNameErr(nil, nil)
	r, err := s.usecase.Login(context.Background(), dtoUser)
	s.Nil(r)
	s.NotNil(err)
}

// User Login Test END


func TestIUserUseCase(t *testing.T) {
	suite.Run(t, &suiteIUserUseCase{})
}
