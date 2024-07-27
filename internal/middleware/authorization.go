package middleware

import (
	"kinolove/internal/service"
	"kinolove/pkg/logger"
	"net/http"
)

type AuthMiddleware struct {
	log         logger.Common
	authService service.AuthService
	renderer    func(w http.ResponseWriter, r *http.Request, servErr *service.ServErr, log logger.Common)
}

func NewAuthMiddleware(authService service.AuthService, log logger.Common,
	renderer func(w http.ResponseWriter, r *http.Request, servErr *service.ServErr, log logger.Common)) *AuthMiddleware {
	return &AuthMiddleware{
		log:         log,
		authService: authService,
		renderer:    renderer,
	}
}

func (a *AuthMiddleware) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := a.authService.VerifyJwt(r)

		if err != nil {
			a.renderer(w, r, err, a.log)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *AuthMiddleware) HasPermission(permId int64, lvl int16) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tok, err := a.authService.VerifyJwt(r)

			if err != nil {
				a.renderer(w, r, err, a.log)
				return
			}

			if err := a.authService.HasPermission(tok, permId, lvl); err != nil {
				a.renderer(w, r, err, a.log)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
