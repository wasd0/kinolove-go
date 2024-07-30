package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/repository"
	"kinolove/internal/service/dto"
	"kinolove/internal/utils/mapper"
	"kinolove/pkg/utils/crypt"
)

type UserServiceImpl struct {
	userRepo    repository.UserRepository
	authService AuthService
}

func NewUserService(repo repository.UserRepository, auth AuthService) *UserServiceImpl {
	return &UserServiceImpl{userRepo: repo, authService: auth}
}

func (u *UserServiceImpl) CreateUser(request dto.UserCreateRequest) (uuid.UUID, *ServErr) {
	isExists, err := u.userRepo.ExistsByUsername(request.Username)

	if err != nil {
		return uuid.Nil, InternalError(err)
	} else if isExists {
		msg := fmt.Sprintf("user with username %s already exists", request.Username)
		return uuid.Nil, BadRequest(errors.New("User exists"), msg)
	}

	hash, cryptErr := crypt.Encode(request.Password)

	if cryptErr != nil {
		return uuid.Nil, InternalError(cryptErr)
	}

	usr := &model.Users{
		Username: request.Username,
		Password: hash,
	}

	err = u.userRepo.Save(usr)
	if err != nil {
		return uuid.Nil, InternalError(err)
	}

	return usr.ID, nil
}

func (u *UserServiceImpl) FindByUsername(username string) (dto.UserSingleResponse, *ServErr) {
	usr, err := u.userRepo.GetByUsername(username)

	if err != nil {
		return dto.UserSingleResponse{}, BadRequest(err, fmt.Sprintf("user with username %s not found", username))
	}

	return mapper.MapUserToSingleResponse(usr), nil
}

func (u *UserServiceImpl) Update(id uuid.UUID, request dto.UserUpdateRequest) *ServErr {
	usr, err := u.userRepo.GetById(id)

	if err != nil {
		return BadRequest(err, fmt.Sprintf("User with id %s not found", id))
	}

	if request.Username != nil {
		isExists, repoErr := u.userRepo.ExistsByUsername(*request.Username)

		if repoErr != nil {
			return InternalError(repoErr)
		} else if isExists && usr.ID != id {
			msg := fmt.Sprintf("user with username %s already exists", *request.Username)
			return BadRequest(errors.New("Already exists"), msg)
		}
	}

	err = mapper.MapUpdateRequestToUser(&request, usr)
	if err != nil {
		return InternalError(err)
	}

	updErr := u.userRepo.Update(usr)

	if updErr != nil {
		return InternalError(updErr)
	}

	return nil
}

func (u *UserServiceImpl) GetByUsername(username string) (*model.Users, *ServErr) {
	usr, err := u.userRepo.GetByUsername(username)

	if err != nil {
		return nil, BadRequest(err, fmt.Sprintf("user with username %s not found", username))
	}

	return usr, nil
}
