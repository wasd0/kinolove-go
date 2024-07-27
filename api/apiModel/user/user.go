package user

import (
	"github.com/pkg/errors"
	"kinolove/internal/service/dto"
	"kinolove/internal/utils/validator"
	"net/http"
)

type ReqUserCreate struct {
	dto.UserCreateRequest
}

func (r ReqUserCreate) Bind(_ *http.Request) error {
	if !validator.ValidateUsername(r.Username) {
		return errors.New("username is too short")
	}

	if !validator.ValidatePassword(r.Password) {
		return errors.New("password is too short")
	}

	return nil
}

type ReqUserUpdate struct {
	dto.UserUpdateRequest
}

func (r ReqUserUpdate) Bind(_ *http.Request) error {
	if r.Username != nil && !validator.ValidateUsername(*r.Username) {
		return errors.New("username is too short")
	}

	if r.Password != nil && !validator.ValidatePassword(*r.Password) {
		return errors.New("password is too short")
	}

	return nil
}
