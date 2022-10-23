package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	mock_helper "github.com/riyanda432/belajar-authentication/mocks/helper"
	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
	repositories "github.com/riyanda432/belajar-authentication/src/domain/repositories/mobile_app"
	infra_model "github.com/riyanda432/belajar-authentication/src/infra/models"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteUserRepository struct {
	suite.Suite
	mockDB       sqlmock.Sqlmock
	db           *gorm.DB
	userRepo     repositories.UserRepository
	user         *entities.User
	userEntities []entities.User
	userModel    infra_model.User
	userColumn   []string

	ctx context.Context
}

func (s *suiteUserRepository) SetupTest() {
	s.ctx = context.Background()
	db, mockDB, err := mock_helper.MockGormDB()
	if err != nil {
		s.T().Errorf(err.Error())
	}
	s.db = db
	s.mockDB = *mockDB

	s.userRepo = NewUserRepository(s.db)
	userEntity := entities.MakeUser(
		uint64(1),
		"Agus Supriyadi",
		"agus@gmail.com",
		"agus123",
		time.Now(),
		time.Now(),
	)
	s.user = userEntity
	s.userEntities = []entities.User{
		*userEntity,
	}
	vm := infra_model.ToUserModel(s.userEntities[0])
	s.userModel = *vm
	vc := s.userModel.GetColumnName()
	s.userColumn = vc
}

func (s *suiteUserRepository) TestPersist_UserPersist() {
	dd := entities.MakeUser(
		1,
		"Agus Supriyadi",
		"agus@gmail.com",
		"agus12356789",
		time.Now(),
		time.Now(),
	)
	e := s.userRepo.Persist(context.Background(), dd)
	s.EqualError(e, "all expectations were already fulfilled, call to database transaction Begin was not expected")
}

func (s *suiteUserRepository) TestPersist_EntityNil() {
	e := s.userRepo.Persist(context.Background(), nil)
	s.EqualError(e, "user entity can not empty")
}

func (s *suiteUserRepository) TestDetailByUserName() {
	s.mockDB.ExpectPrepare(
		`SELECT \"users\".\"id\",\"users\".\"full_name\",\"users\".\"user_name\",\"users\".\"password\",\"users\".\"created_at\",\"users\".\"updated_at\" FROM \"users\" WHERE \"users\".\"user_name\" \= \$1 ORDER BY \"users\".\"id\" LIMIT 1`,
	).ExpectQuery().WithArgs(
		sqlmock.AnyArg(),
	).WillReturnRows(
		sqlmock.NewRows(
			s.userColumn,
		).AddRow(
			s.userModel.ID,
			s.userModel.FullName,
			s.userModel.UserName,
			s.userModel.Password,
			s.userModel.CreatedAt,
			s.userModel.UpdatedAt,
		),
	)

	expectV := s.userEntities[0]

	se, err := s.userRepo.DetailByUserName(s.ctx, "agus@gmail.com")
	s.Equal(expectV, *se)
	s.Nil(err)
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, &suiteUserRepository{})
}
