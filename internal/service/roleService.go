package service

import (
	"github.com/go-jet/jet/v2/qrm"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"kinolove/internal/repository"
)

type RoleServiceImpl struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) *RoleServiceImpl {
	return &RoleServiceImpl{roleRepo: roleRepo}
}

func (r *RoleServiceImpl) GetUserRolesIds(usrId uuid.UUID) (*[]int64, *ServErr) {
	ids, err := r.roleRepo.GetUserRolesIds(usrId)

	if err != nil && !errors.As(err, qrm.ErrNoRows) {
		return nil, InternalError(err)
	}

	return ids, nil
}
