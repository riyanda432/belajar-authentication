package transformer_v1_mobile

import (
	"time"

	entities "github.com/riyanda432/belajar-authentication/src/domain/entities"
	"github.com/riyanda432/belajar-authentication/src/infra/constants"
)

type UserTransformCreateUpdate struct {
	ID        uint64 `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func TransformCreateUpdate(o *entities.User) *UserTransformCreateUpdate {
	if o == nil {
		return nil
	}
	return &UserTransformCreateUpdate{
		ID:        o.GetID(),
		CreatedAt: o.GetCreatedAt().Format(constants.ISODateTimeFormat),
		UpdatedAt: o.GetUpdatedAt().Format(constants.ISODateTimeFormat),
	}
}

type UserTransformSuccessLogin struct {
	Username string `json:"username"`
	LoginAt  string `json:"loginAt"`
}

func TransformSuccessLogin(o *entities.User) *UserTransformSuccessLogin {
	if o == nil {
		return nil
	}
	return &UserTransformSuccessLogin{
		Username: o.GetUsername(),
		LoginAt:  time.Now().Format(constants.ISODateTimeFormat),
	}
}
