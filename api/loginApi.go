package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"kinolove/api/apiModel"
	"kinolove/api/apiModel/login"
	"kinolove/internal/middleware"
	"kinolove/internal/service"
	"kinolove/pkg/logger"
	"net/http"
)

type LoginApi struct {
	loginService service.LoginService
	log          logger.Common
	auth         *middleware.AuthMiddleware
}

func NewLoginApi(log logger.Common, loginService service.LoginService, auth *middleware.AuthMiddleware) *LoginApi {
	return &LoginApi{loginService, log, auth}
}

func (l *LoginApi) Register() (string, func(router chi.Router)) {
	return "/api/v1", l.Handle
}

func (l *LoginApi) Handle(router chi.Router) {
	router.Post("/login", l.Login)
	router.With(l.auth.Authenticator).Post("/logout", l.Logout)
}

func (l *LoginApi) Login(w http.ResponseWriter, r *http.Request) {
	request := login.ReqLogin{}

	if err := render.Bind(r, &request); err != nil {
		RenderError(w, r, service.BadRequest(err, "Failed get request body"), l.log)
		return
	}

	if jwt, err := l.loginService.Login(w, request.LoginRequest); err != nil {
		RenderError(w, r, err, l.log)
		return
	} else {
		response := apiModel.RestResponse[string]{Data: &jwt}
		if renderErr := render.Render(w, r, &response); renderErr != nil {
			RenderError(w, r, err, l.log)
		}

	}
}

func (l *LoginApi) Logout(w http.ResponseWriter, r *http.Request) {
	err := l.loginService.Logout(w)

	if err != nil {
		RenderError(w, r, err, l.log)
		return
	}
}
