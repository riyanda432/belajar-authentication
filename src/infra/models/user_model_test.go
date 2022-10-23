package infra_model

import (
	"testing"
	"time"

	domain_models "github.com/riyanda432/belajar-authentication/src/domain/entities"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite

	uModel  User
	uEntity *domain_models.User
}

func (s *UserSuite) SetupTest() {
	s.uEntity = MockUserModel()
}

var MockUserModel = func() *domain_models.User {
	ss:= domain_models.MakeUser(
		1,                       //id
		"SRV-000000002",         //code
		"BUDI-JAYA",             //name
		"Jalan Setiabudi no 67", //address
		time.Now(),              //createdAt
		time.Now(),              //updatedAt
	)
	return ss
}

func (s *UserSuite) TestGetColumnName() {
	r := s.uModel.GetColumnName()
	s.NotNil(r)
}

func (s *UserSuite) TestMakeUser() {
	user := User{
		ID: uint64(1),
		FullName: "Agus Supriyadi",
		UserName: "agus@gmail.com",
		Password: "agus123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	ts := user.ToEntity()
	s.NotNil(ts)
}

func (s *UserSuite) TestToUserModel() {
	user := User{
		ID: uint64(1),
		FullName: "Agus Supriyadi",
		UserName: "agus@gmail.com",
		Password: "agus123456789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	ts := user.ToEntity()
	ss := ToUserModel(*ts)
	s.NotNil(ss)
}

func TestUser(t *testing.T) {
	suite.Run(t, &UserSuite{})
}
