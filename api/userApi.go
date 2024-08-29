package api

import (
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"kinolove/api/apiModel"
	"kinolove/api/apiModel/user"
	"kinolove/internal/middleware"
	"kinolove/internal/service"
	"kinolove/internal/service/dto"
	"kinolove/pkg/constants/perms"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserApi struct {
	userService service.UserService
	auth        *middleware.AuthMiddleware
	authService service.AuthService
}

func NewUserApi(userService service.UserService,
	auth *middleware.AuthMiddleware, authService service.AuthService) *UserApi {
	return &UserApi{userService: userService, auth: auth, authService: authService}
}

func (u *UserApi) Register() (string, func(router chi.Router)) {
	return "/api/v1/users", u.Handle
}

func (u *UserApi) Handle(router chi.Router) {
	router.Post("/", u.createUser)
	router.With(u.auth.Authenticator).Get("/{username}", u.findByUsername)
	router.With(u.auth.HasPermission(perms.User, perms.Edit)).Put("/{id}", u.update)
}

func (u *UserApi) createUser(w http.ResponseWriter, r *http.Request) {
	request := user.ReqUserCreate{}

	if err := render.Bind(r, &request); err != nil {
		RenderError(w, r, service.BadRequest(err, "Failed get request body"))
		return
	}

	if id, err := u.userService.CreateUser(request.UserCreateRequest); err != nil {
		RenderError(w, r, err)
	} else {
		response := apiModel.RestResponse[uuid.UUID]{Data: &id}
		if renderErr := render.Render(w, r, &response); renderErr != nil {
			RenderError(w, r, service.InternalError(renderErr))
			return
		}
	}
}

func (u *UserApi) findByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	byUsername, err := u.userService.FindByUsername(username)
	if err != nil {
		RenderError(w, r, err)
		return
	}

	response := apiModel.RestResponse[dto.UserSingleResponse]{Data: &byUsername}
	if renderErr := render.Render(w, r, &response); renderErr != nil {
		RenderError(w, r, err)
		return
	}
}

func (u *UserApi) update(w http.ResponseWriter, r *http.Request) {
	uuidStr := chi.URLParam(r, "id")

	if err := uuid.Validate(uuidStr); err != nil {
		RenderError(w, r, service.BadRequest(errors.New("Wrong id"), "wrong user id"))
		return
	}

	id, err := uuid.Parse(uuidStr)

	if err != nil {
		RenderError(w, r, service.InternalError(err))
		return
	}

	if jwt, _, err := jwtauth.FromContext(r.Context()); err != nil {
		RenderError(w, r, service.Unauthorized(err))
		return
	} else if u.authService.IsAuthenticated(&jwt, id) != nil {
		RenderError(w, r, service.Forbidden(err))
		return
	}

	request := user.ReqUserUpdate{}
	if err = render.Bind(r, &request); err != nil {
		RenderError(w, r, service.BadRequest(err, "Failed get request body"))
		return
	}

	servErr := u.userService.Update(id, request.UserUpdateRequest)

	if servErr != nil {
		RenderError(w, r, servErr)
		return
	}
}
