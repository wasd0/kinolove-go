package login

import (
	"github.com/pkg/errors"
	"kinolove/internal/service/dto"
	"kinolove/internal/utils/validator"
	"net/http"
)

type ReqLogin struct {
	dto.LoginRequest
}

func (r2 *ReqLogin) Bind(r *http.Request) error {
	if !validator.ValidateUsername(r2.Username) {
		return errors.New("Invalid Username  length")
	}

	if !validator.ValidatePassword(r2.Password) {
		return errors.New("Invalid Password  length")
	}

	return nil
}
