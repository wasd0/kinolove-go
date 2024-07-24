package service

import (
	"fmt"
	"github.com/google/uuid"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/repository"
	"kinolove/internal/service/dto"
	"kinolove/internal/utils/mapper"
	"kinolove/pkg/utils/crypt"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: repo}
}

func (u *UserServiceImpl) CreateUser(request dto.UserCreateRequest) (uuid.UUID, error) {
	isExists, err := u.userRepo.ExistsByUsername(request.Username)

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

	err = u.userRepo.Save(usr)
	if err != nil {
		return getCreationErr("errorUtils while saving new user: %v", err)
	}

	return usr.ID, nil
}

func (u *UserServiceImpl) FindByUsername(username string) (dto.UserSingleResponse, error) {
	usr, err := u.userRepo.GetByUsername(username)

	if err != nil {
		return dto.UserSingleResponse{}, err
	}

	return mapper.MapUserToSingleResponse(usr), nil
}

func (u *UserServiceImpl) Update(id uuid.UUID, request dto.UserUpdateRequest) error {
	if request.Username != nil {
		isExists, err := u.userRepo.ExistsByUsername(*request.Username)

		if err != nil {
			return fmt.Errorf("eror while checking existence of user '%s': %v", *request.Username, err)
		} else if isExists {
			return fmt.Errorf("user with username '%s' already exists", *request.Username)
		}
	}

	usr, err := u.userRepo.GetById(id)

	if err != nil {
		return getUpdateErr(id, err)
	}

	err = mapper.MapUpdateRequestToUser(&request, usr)
	if err != nil {
		return err
	}

	updErr := u.userRepo.Update(usr)

	if updErr != nil {
		return getUpdateErr(id, err)
	}

	return nil
}

func getCreationErr(format string, args interface{}) (uuid.UUID, error) {
	return uuid.Nil, fmt.Errorf(format, args)
}

func getUpdateErr(id uuid.UUID, err error) error {
	return fmt.Errorf("error while getting user '%s': %v\"", id.String(), err)
}
