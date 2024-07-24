package user

import (
	"github.com/pkg/errors"
	"kinolove/internal/service/dto"
	"net/http"
)

type ReqUserCreate struct {
	dto.UserCreateRequest
}

func (r ReqUserCreate) Bind(_ *http.Request) error {
	if len(r.Username) < 5 {
		return errors.New("username is too short")
	}

	if len(r.Password) < 8 {
		return errors.New("password is too short")
	}

	return nil
}

type ReqUserUpdate struct {
	dto.UserUpdateRequest
}

func (r ReqUserUpdate) Bind(_ *http.Request) error {
	if r.Username != nil && len(*r.Username) < 5 {
		return errors.New("username is too short")
	}

	if r.Password != nil && len(*r.Password) < 8 {
		return errors.New("password is too short")
	}

	return nil
}
