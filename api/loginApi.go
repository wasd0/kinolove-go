package api

import (
	"github.com/go-chi/chi/v5"
	"kinolove/api/apiModel/login"
	"kinolove/internal/service"
	"kinolove/pkg/logger"
	"net/http"
)

type LoginApi struct {
	loginService service.LoginService
	log          logger.Common
}

func NewLoginApi(log logger.Common, loginService service.LoginService) *LoginApi {
	return &LoginApi{loginService, log}
}

func (l *LoginApi) Register() (string, func(router chi.Router)) {
	return "/api/v1", l.Handle
}

func (l *LoginApi) Handle(router chi.Router) {
	router.Post("/login", l.Login)
	router.Post("/logout", l.Logout)
}

func (l *LoginApi) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := login.ReqLogin{}
	if username, password, isOk := r.BasicAuth(); isOk {
		loginRequest.Username = username
		loginRequest.Password = password
	}

	if err := l.loginService.Login(w, loginRequest.LoginRequest); err != nil {
		RenderError(w, r, err, l.log)
		return
	}
}

func (l *LoginApi) Logout(w http.ResponseWriter, r *http.Request) {
	err := l.loginService.Logout(w, r)

	if err != nil {
		RenderError(w, r, err, l.log)
		return
	}
}
