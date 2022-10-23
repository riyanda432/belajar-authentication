package data_transfer_object_mobile_app_v1

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	error_helper "github.com/riyanda432/belajar-authentication/src/helper"
	infra_error "github.com/riyanda432/belajar-authentication/src/infra/errors"

	"net/mail"
)

type IUserDTO interface {
	Validate() error
}

type UserCreate struct {
	FullName        string
	UserName        string
	Password        string
	ConfirmPassword string
}

func NewUserCreateDTO(
	fn string,
	un string,
	p string,
	cp string,
) IUserDTO {
	return &UserCreate{
		FullName:        fn,
		UserName:        un,
		Password:        p,
		ConfirmPassword: cp,
	}
}

func (dto *UserCreate) Validate() error {
	infra_error.InitErrorDicts()

	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.FullName, validation.Required, validation.Length(2, 25)),
		validation.Field(&dto.UserName, validation.Required, validation.By(isValidEmail)),
		validation.Field(&dto.Password, validation.Required, validation.Length(12, 100)),
		validation.Field(&dto.ConfirmPassword, validation.Required, validation.In(dto.Password)),
	); err != nil {
		retErr := error_helper.NewCommonError(infra_error.INVALID_REQUEST_CREATE_USER, err)
		retErr.SetValidationMessage(err)
		return retErr
	}

	return nil
}

type UserLogin struct {
	UserName        string
	Password        string
}

func NewUserLoginDTO(
	un string,
	p string,
) IUserDTO {
	return &UserLogin{
		UserName:        un,
		Password:        p,
	}
}

func (dto *UserLogin) Validate() error {
	infra_error.InitErrorDicts()

	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserName, validation.Required),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		retErr := error_helper.NewCommonError(infra_error.INVALID_REQUEST_CREATE_USER, err)
		retErr.SetValidationMessage(err)
		return retErr
	}

	return nil
}

func isValidEmail(value interface{}) error {
	_, err := mail.ParseAddress(value.(string))
	if err != nil {
		return errors.New("Please provide a valid email address")
	}
	return nil
}
