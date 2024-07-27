package service

import (
	"kinolove/internal/service/dto"
	"net/http"
)

type LoginServiceImpl struct {
	userService       UserService
	authService       AuthService
	permissionService PermissionService
}

func NewLoginService(us UserService,
	authServ AuthService,
	permService PermissionService) *LoginServiceImpl {
	return &LoginServiceImpl{
		userService:       us,
		authService:       authServ,
		permissionService: permService,
	}
}

func (l *LoginServiceImpl) Login(w http.ResponseWriter, request dto.LoginRequest) *ServErr {
	usr, userErr := l.userService.GetByUsername(request.Username)

	if userErr != nil {
		return userErr
	}

	if authErr := l.authService.Authenticate(usr, request.Password); authErr != nil {
		return authErr
	}

	perms, permErr := l.permissionService.GetAllUserPermissions(usr)

	if permErr != nil {
		return permErr
	}

	jwtToken, authErr := l.authService.GetJwtToken(usr.ID, perms)

	if authErr != nil {
		return authErr
	}

	w.Header().Add("Set-Cookie", "jwt="+jwtToken)

	return nil
}

func (l *LoginServiceImpl) Logout(w http.ResponseWriter) *ServErr {
	w.Header().Set("Set-Cookie", "jwt=")

	return nil
}
