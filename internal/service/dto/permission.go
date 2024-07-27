package dto

import . "kinolove/internal/entity/.gen/kinolove/public/model"

type AllUserPermission struct {
	UserPerms *[]UsersPermissions
	RolePerms *[]RolesPermissions
}
