package apiProvider

import (
	"kinolove/api"
	"kinolove/internal/app/serviceProvider"
	"kinolove/internal/middleware"
	"kinolove/pkg/logger"
)

type ApiProvider struct {
	serviceProvider *serviceProvider.ServiceProvider
	log             logger.Common

	defaultApi *api.DefaultApi
	user       *api.UserApi
	movie      *api.MovieApi
	login      *api.LoginApi

	authMid *middleware.AuthMiddleware
}

func InitApi(serviceProvider *serviceProvider.ServiceProvider, log logger.Common) *ApiProvider {
	auth := middleware.NewAuthMiddleware(serviceProvider.AuthService(), log, api.RenderError)
	return &ApiProvider{serviceProvider: serviceProvider, log: log, authMid: auth}
}

func (ap *ApiProvider) DefaultApi() *api.DefaultApi {
	if ap.defaultApi != nil {
		return ap.defaultApi
	}

	dApi := api.NewDefaultApi(ap.log)
	ap.defaultApi = dApi
	return ap.defaultApi
}

func (ap *ApiProvider) UserApi() *api.UserApi {
	if ap.user != nil {
		return ap.user
	}

	dApi := api.NewUserApi(ap.log, ap.serviceProvider.UserService(), ap.authMid)
	ap.user = dApi
	return ap.user
}

func (ap *ApiProvider) MovieApi() *api.MovieApi {
	if ap.movie != nil {
		return ap.movie
	}

	dApi := api.NewMovieApi(ap.log, ap.serviceProvider.MovieService())
	ap.movie = dApi
	return ap.movie
}

func (ap *ApiProvider) LoginApi() *api.LoginApi {
	if ap.login != nil {
		return ap.login
	}

	dApi := api.NewLoginApi(ap.log, ap.serviceProvider.LoginService())
	ap.login = dApi
	return ap.login
}
