package service

import (
	"errors"
	"github.com/go-jet/jet/v2/qrm"
	. "kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/repository"
	"kinolove/internal/service/dto"
)

type PermissionServiceImpl struct {
	permRepo    repository.PermissionRepository
	roleService RoleService
}

func NewPermissionService(permRepo repository.PermissionRepository, roleService RoleService) *PermissionServiceImpl {
	return &PermissionServiceImpl{
		permRepo:    permRepo,
		roleService: roleService,
	}
}

func (p *PermissionServiceImpl) GetAllUserPermissions(usr *Users) (*dto.AllUserPermission, *ServErr) {
	userRoleIds, rolesErr := p.roleService.GetUserRolesIds(usr.ID)

	if rolesErr != nil {
		return nil, rolesErr
	}

	rolesPerms, permErr := p.permRepo.FindRolePermissions(userRoleIds)

	if permErr != nil && !errors.Is(permErr, qrm.ErrNoRows) {
		return nil, InternalError(permErr)
	}

	userPerms, permErr := p.permRepo.FindUserPermissions(usr.ID)

	if permErr != nil && !errors.Is(permErr, qrm.ErrNoRows) {
		return nil, InternalError(permErr)
	}

	return &dto.AllUserPermission{
		UserPerms: userPerms,
		RolePerms: rolesPerms,
	}, nil
}
