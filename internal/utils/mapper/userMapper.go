package mapper

import (
	"kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/service/dto"
	"kinolove/pkg/utils/clock"
	"kinolove/pkg/utils/crypt"
	"time"
)

func MapUserToSingleResponse(usr *model.Users) dto.UserSingleResponse {
	if usr == nil {
		return dto.UserSingleResponse{}
	}

	dateReg := clock.GetUtc(usr.DateReg)
	datePwdUpd := clock.GetUtc(usr.DatePassUpd)

	return dto.UserSingleResponse{
		Username:    usr.Username,
		IsActive:    usr.IsActive,
		DateReg:     dateReg,
		DatePassUpd: datePwdUpd,
	}
}

func MapUpdateRequestToUser(request *dto.UserUpdateRequest, user *model.Users) error {
	if request == nil {
		return nil
	}

	if request.Username != nil {
		user.Username = *request.Username
	}

	if request.Password != nil {
		hash, err := crypt.Encode(*request.Password)

		if err != nil {
			return err
		}

		user.Password = hash
		now := time.Now()
		user.DatePassUpd = clock.GetUtc(&now)
	}

	return nil
}
