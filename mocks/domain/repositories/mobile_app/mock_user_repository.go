package mock_domain

import (
	"context"

	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
	repos "github.com/riyanda432/belajar-authentication/src/domain/repositories/mobile_app"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{}
}

// implement interface
var _ repos.UserRepository = &MockUserRepository{}

func (m *MockUserRepository) MockPersist(res *entities.User, err error) {
	m.Mock.On("Persist", mock.Anything, mock.Anything).Return(res, nil)
}

func (m *MockUserRepository) MockPersistErr(res *entities.User, err error) {
	m.Mock.On("Persist", mock.Anything, mock.Anything).Return(nil, err)
}

// Persist implements domain_repositories.UserRepository
func (m *MockUserRepository) Persist(context context.Context, user *entities.User) error {
	args := m.Called(context, user)

	var err error

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return err
}

func (m *MockUserRepository) MockDetailByUserNameSuccess(res *entities.User, err error) {
	m.Mock.On("DetailByUserName", mock.Anything, mock.Anything).Return(res, nil)
}

func (m *MockUserRepository) MockDetailByUserNameErr(res *entities.User, err error) {
	m.Mock.On("DetailByUserName", mock.Anything, mock.Anything).Return(nil, err)
}

// DetailByUserName implements domain_repositories.UserRepository
func (m *MockUserRepository) DetailByUserName(context context.Context, userName string) (*entities.User, error) {
	args := m.Called(context, userName)

	var err error
	var user *entities.User
	if n, ok := args.Get(0).(*entities.User); ok {
		user = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return user, err
}
