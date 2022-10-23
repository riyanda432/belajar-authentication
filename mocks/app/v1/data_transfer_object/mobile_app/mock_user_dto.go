package mock_dto_v1_mobile_app

import (
	dto "github.com/riyanda432/belajar-authentication/src/app/v1/data_transfer_object/mobile_app"
	"github.com/stretchr/testify/mock"
)

type MockUserDTO struct {
	mock.Mock
}

func (m *MockUserDTO) Validate() error {
	args := m.Called()
	var err error

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return err
}

func (m *MockUserDTO) MakeMockUserCreate() *dto.UserCreate {
	return &dto.UserCreate{
		FullName:        "Agus Supriyadi",
		UserName:        "agus@gmail.com",
		Password:        "agus1234567890",
		ConfirmPassword: "agus1234567890",
	}
}

func (m *MockUserDTO) MakeMockUserLogin() *dto.UserLogin {
	return &dto.UserLogin{
		UserName: "agus@gmail.com",
		Password: "agus12345678912",
	}
}

var _ dto.IUserDTO = &MockUserDTO{}
