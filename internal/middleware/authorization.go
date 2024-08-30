package middleware

import (
	"kinolove/internal/service"
	"net/http"
)

type AuthMiddleware struct {
	authService service.AuthService
	renderer    func(w http.ResponseWriter, r *http.Request, servErr *service.ServErr)
}

func NewAuthMiddleware(authService service.AuthService,
	renderer func(w http.ResponseWriter, r *http.Request, servErr *service.ServErr)) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
		renderer:    renderer,
	}
}

func (a *AuthMiddleware) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := a.authService.VerifyJwt(r)

		if err != nil {
			a.renderer(w, r, err)
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
				a.renderer(w, r, err)
				return
			}

			if err := a.authService.HasPermission(tok, permId, lvl); err != nil {
				a.renderer(w, r, err)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
