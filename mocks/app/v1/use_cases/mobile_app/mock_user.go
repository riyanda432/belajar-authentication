package mock_usecases

import (
	"context"

	dto "github.com/riyanda432/belajar-authentication/src/app/v1/data_transfer_object/mobile_app"
	usecases "github.com/riyanda432/belajar-authentication/src/app/v1/use_cases/mobile_app"
	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockIUserUseCase struct {
	mock.Mock
}

func (m *MockIUserUseCase) MockCreateError(res *entities.User, err error) {
	m.Mock.On("Create", mock.Anything, mock.Anything).Return(nil, err)
}

func (m *MockIUserUseCase) MockCreateSuccess(res *entities.User, err error) {
	m.Mock.On("Create", mock.Anything, mock.Anything).Return(res, nil)
}

// Create implements usecases.IUserUseCase
func (m *MockIUserUseCase) Create(ctx context.Context, do dto.IUserDTO) (*entities.User, error) {
	args := m.Called(ctx, do)

	var res *entities.User
	var err error

	if n, ok := args.Get(0).(*entities.User); ok {
		res = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return res, err
}

func (m *MockIUserUseCase) MockLoginError(res *entities.User, err error) {
	m.Mock.On("Login", mock.Anything, mock.Anything).Return(nil, err)
}

func (m *MockIUserUseCase) MockLoginSuccess(res *entities.User, err error) {
	m.Mock.On("Login", mock.Anything, mock.Anything).Return(res, nil)
}

// Login implements usecases.IUserUseCase
func (m *MockIUserUseCase) Login(ctx context.Context, do dto.IUserDTO) (*entities.User, error) {
	args := m.Called(ctx, do)

	var res *entities.User
	var err error

	if n, ok := args.Get(0).(*entities.User); ok {
		res = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return res, err
}


var _ usecases.IUserUseCase = &MockIUserUseCase{}
