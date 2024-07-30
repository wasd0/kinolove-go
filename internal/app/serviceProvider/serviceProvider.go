package serviceProvider

import (
	"kinolove/internal/app/repoProvider"
	"kinolove/internal/service"
	"kinolove/pkg/utils/jwtUtils"
)

type ServiceProvider struct {
	provider *repoProvider.RepoProvider

	jwt   *jwtUtils.Auth
	user  service.UserService
	movie service.MovieService
	login service.LoginService
	auth  service.AuthService
	role  service.RoleService
	perm  service.PermissionService
}

func InitServices(provider *repoProvider.RepoProvider, jwt *jwtUtils.Auth) *ServiceProvider {
	return &ServiceProvider{provider: provider, jwt: jwt}
}

func (sp *ServiceProvider) UserService() service.UserService {
	if sp.user != nil {
		return sp.user
	}

	user := service.NewUserService(sp.provider.UserRepo(), sp.AuthService())
	sp.user = user
	return sp.user
}

func (sp *ServiceProvider) MovieService() service.MovieService {
	if sp.movie != nil {
		return sp.movie
	}

	movie := service.NewMovieService(sp.provider.MovieRepo())
	sp.movie = movie
	return sp.movie
}

func (sp *ServiceProvider) LoginService() service.LoginService {
	if sp.login != nil {
		return sp.login
	}

	login := service.NewLoginService(sp.UserService(), sp.AuthService(), sp.PermissionService())
	sp.login = login
	return sp.login
}

func (sp *ServiceProvider) AuthService() service.AuthService {
	if sp.auth != nil {
		return sp.auth
	}

	auth := service.NewAuthService(sp.jwt)
	sp.auth = auth
	return sp.auth
}

func (sp *ServiceProvider) RoleService() service.RoleService {
	if sp.role != nil {
		return sp.role
	}

	role := service.NewRoleService(sp.provider.RoleRepo())
	sp.role = role
	return sp.role
}

func (sp *ServiceProvider) PermissionService() service.PermissionService {
	if sp.perm != nil {
		return sp.perm
	}

	perm := service.NewPermissionService(sp.provider.PermissionRepo(), sp.RoleService())
	sp.perm = perm
	return sp.perm
}
