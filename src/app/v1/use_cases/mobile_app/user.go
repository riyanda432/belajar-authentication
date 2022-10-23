package usecases_mobile_app_v1

import (
	"context"
	"errors"

	error_helper "github.com/riyanda432/belajar-authentication/src/helper"
	infra_error "github.com/riyanda432/belajar-authentication/src/infra/errors"

	dto "github.com/riyanda432/belajar-authentication/src/app/v1/data_transfer_object/mobile_app"
	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
	repositories "github.com/riyanda432/belajar-authentication/src/domain/repositories/mobile_app"
	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	Create(ctx context.Context, do dto.IUserDTO) (*entities.User, error)
	Login(ctx context.Context, do dto.IUserDTO) (*entities.User, error)
}

type userUseCase struct {
	repository repositories.UserRepository
}

func NewUserUseCase(or repositories.UserRepository) IUserUseCase {
	return &userUseCase{
		repository: or,
	}
}

func (u *userUseCase) Create(ctx context.Context, do dto.IUserDTO) (*entities.User, error) {
	userCreate, ok := do.(*dto.UserCreate)
	if !ok {
		return nil, errors.New("type assertion failed to dto.UserCreate")
	}
	var user *entities.User
	// check if user has been registered
	result, err := u.repository.DetailByUserName(ctx, userCreate.UserName)
	if err != nil {
		return nil, error_helper.NewCommonError(infra_error.FAILED_RETRIEVE_USER, err)
	}
	// if user has been registered return error
	if result != nil {
		return nil, error_helper.NewCommonError(infra_error.INVALID_USER, err)
	}
	// encrypt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), 12)
	if err != nil {
		return nil, err
	}
	passString := string(hashedPassword)

	// insert into users (Postgres)
	user = entities.CreateUser(
		userCreate.FullName,
		userCreate.UserName,
		passString,
	)

	errPersistVoucher := u.repository.Persist(ctx, user)
	if errPersistVoucher != nil {
		return nil, errPersistVoucher
	}

	return user, nil
}

func (u *userUseCase) Login(ctx context.Context, do dto.IUserDTO) (*entities.User, error) {
	userLogin, ok := do.(*dto.UserLogin)
	if !ok {
		return nil, errors.New("type assertion failed to dto.UserLogin")
	}
	result, err := u.repository.DetailByUserName(ctx, userLogin.UserName)
	if err != nil {
		return nil, error_helper.NewCommonError(infra_error.FAILED_RETRIEVE_USER, err)
	}

	if result == nil {
		return nil, error_helper.NewCommonError(infra_error.USER_NOT_FOUND, err)
	}
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(result.GetPassword()), []byte(userLogin.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		return nil, error_helper.NewCommonError(infra_error.USER_WRONG_PASSWORD, err)
	}

	return result, nil
}
