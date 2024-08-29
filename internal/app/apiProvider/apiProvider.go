package apiProvider

import (
	"kinolove/api"
	"kinolove/internal/app/serviceProvider"
	"kinolove/internal/middleware"
)

type ApiProvider struct {
	serviceProvider *serviceProvider.ServiceProvider

	defaultApi *api.DefaultApi
	user       *api.UserApi
	movie      *api.MovieApi
	login      *api.LoginApi

	authMid *middleware.AuthMiddleware
}

func InitApi(serviceProvider *serviceProvider.ServiceProvider) *ApiProvider {
	auth := middleware.NewAuthMiddleware(serviceProvider.AuthService(), api.RenderError)
	return &ApiProvider{serviceProvider: serviceProvider, authMid: auth}
}

func (ap *ApiProvider) DefaultApi() *api.DefaultApi {
	if ap.defaultApi != nil {
		return ap.defaultApi
	}

	dApi := api.NewDefaultApi()
	ap.defaultApi = dApi
	return ap.defaultApi
}

func (ap *ApiProvider) UserApi() *api.UserApi {
	if ap.user != nil {
		return ap.user
	}

	dApi := api.NewUserApi(ap.serviceProvider.UserService(), ap.authMid, ap.serviceProvider.AuthService())
	ap.user = dApi
	return ap.user
}

func (ap *ApiProvider) MovieApi() *api.MovieApi {
	if ap.movie != nil {
		return ap.movie
	}

	dApi := api.NewMovieApi(ap.serviceProvider.MovieService(), ap.authMid)
	ap.movie = dApi
	return ap.movie
}

func (ap *ApiProvider) LoginApi() *api.LoginApi {
	if ap.login != nil {
		return ap.login
	}

	dApi := api.NewLoginApi(ap.serviceProvider.LoginService(), ap.authMid)
	ap.login = dApi
	return ap.login
}
