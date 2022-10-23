package request_v1_mobile

import (
	"encoding/json"
	"fmt"
	dto "github.com/riyanda432/belajar-authentication/src/app/v1/data_transfer_object/mobile_app"
	error_helper "github.com/riyanda432/belajar-authentication/src/helper"
	common_error "github.com/riyanda432/belajar-authentication/src/infra/errors"
	"net/http"
)

// IUserRequest ...
type IUserRequest interface {
	Validate(r *http.Request) (dto.IUserDTO, error)
}

type UserCreate struct {
	FullName        string
	UserName        string
	Password        string
	ConfirmPassword string
}

func (req *UserCreate) Validate(r *http.Request) (dto.IUserDTO, error) {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common_error.InitErrorDicts()
		return nil, error_helper.NewCommonError(
			common_error.INVALID_PAYLOAD_CREATE_USER,
			fmt.Errorf("invalid JSON format"),
		)
	}

	d := dto.NewUserCreateDTO(
		req.FullName,
		req.UserName,
		req.Password,
		req.ConfirmPassword,
	)

	if err := d.Validate(); err != nil {
		return nil, err
	}

	return d, nil
}

type UserLogin struct {
	UserName string
	Password string
}

func (req *UserLogin) Validate(r *http.Request) (dto.IUserDTO, error) {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		common_error.InitErrorDicts()
		return nil, error_helper.NewCommonError(
			common_error.INVALID_PAYLOAD_LOGIN_USER,
			fmt.Errorf("invalid JSON format"),
		)
	}

	d := dto.NewUserLoginDTO(
		req.UserName,
		req.Password,
	)

	if err := d.Validate(); err != nil {
		return nil, err
	}

	return d, nil
}
