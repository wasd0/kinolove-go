package mapper

import (
	"kinolove/internal/service/dto"
	"kinolove/pkg/utils/jwt"
)

func PermissionToJwt(perms *dto.AllUserPermission) (*map[int64]jwt.Permission, *map[int64]jwt.Permission) {
	var rolePerms, userPerms map[int64]jwt.Permission

	if perms == nil {
		return nil, nil
	}

	if perms.UserPerms != nil {
		userPerms = make(map[int64]jwt.Permission, len(*perms.UserPerms))

		for _, perm := range *perms.UserPerms {
			userPerms[perm.PermissionID] = jwt.Permission{
				TargetLvl: perm.TargetLevel,
				GlobalLvl: perm.GlobalLevel,
			}
		}

	}

	if perms.RolePerms != nil {
		rolePerms = make(map[int64]jwt.Permission, len(*perms.RolePerms))

		for _, perm := range *perms.RolePerms {
			rolePerms[perm.PermissionID] = jwt.Permission{
				TargetLvl: perm.TargetLevel,
				GlobalLvl: perm.GlobalLevel,
			}
		}
	}

	return &userPerms, &rolePerms
}
