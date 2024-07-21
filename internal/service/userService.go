package service

import (
	"fmt"
	"github.com/google/uuid"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/repository"
	"kinolove/internal/service/dto/user"
	"kinolove/pkg/utils/clock"
	"kinolove/pkg/utils/crypt"
)

type UserServiceImpl struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: &repo}
}

func (u *UserServiceImpl) CreateUser(request user.CreateRequest) (uuid.UUID, error) {
	repo := *u.userRepo
	isExists, err := repo.ExistsByUsername(request.Username)

	if err != nil {
		return uuid.Nil, err
	} else if isExists {
		return getCreationErr("user with username '%s' already exists", request.Username)
	}

	hash, cryptErr := crypt.Encode(request.Password)

	if cryptErr != nil {
		return getCreationErr("failed generate password hash - bcrypt errorUtils: %v", cryptErr)
	}

	usr := &model.Users{
		Username: request.Username,
		Password: hash,
	}

	err = repo.Save(usr)
	if err != nil {
		return getCreationErr("errorUtils while saving new user: %v", err)
	}

	return usr.ID, nil
}

func (u *UserServiceImpl) FindByUsername(username string) (user.SingleResponse, error) {
	repo := *u.userRepo
	usr, err := repo.GetByUsername(username)

	if err != nil {
		return user.SingleResponse{}, err
	}

	return mapUserToSingleResponse(usr), nil
}

func getCreationErr(format string, args interface{}) (uuid.UUID, error) {
	return uuid.Nil, fmt.Errorf(format, args)
}

func mapUserToSingleResponse(usr *model.Users) user.SingleResponse {
	if usr == nil {
		return user.SingleResponse{}
	}

	dateReg := clock.GetUtc(usr.DateReg)
	datePwdUpd := clock.GetUtc(usr.DatePassUpd)

	return user.SingleResponse{
		Username:    usr.Username,
		IsActive:    usr.IsActive,
		DateReg:     dateReg,
		DatePassUpd: datePwdUpd,
	}
}
