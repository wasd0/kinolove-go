package mapper

import (
	"kinolove/internal/service/dto"
)

func PermissionToMaps(perms *dto.AllUserPermission) (*map[int64]int16, *map[int64]int16) {
	var rolePerms, userPerms map[int64]int16

	if perms == nil {
		u := make(map[int64]int16)
		r := make(map[int64]int16)
		return &u, &r
	}

	if perms.UserPerms != nil {
		userPerms = make(map[int64]int16, len(*perms.UserPerms))

		for _, perm := range *perms.UserPerms {
			userPerms[perm.PermissionID] = perm.Level
		}

	} else {
		userPerms = make(map[int64]int16)
	}

	if perms.RolePerms != nil {
		rolePerms = make(map[int64]int16, len(*perms.RolePerms))

		for _, perm := range *perms.RolePerms {
			rolePerms[perm.PermissionID] = perm.Level
		}
	} else {
		rolePerms = make(map[int64]int16)
	}

	return &userPerms, &rolePerms
}
