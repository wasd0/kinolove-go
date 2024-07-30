package service

import (
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwt"
	. "kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/service/dto"
	"net/http"
)

type UserService interface {
	CreateUser(request dto.UserCreateRequest) (uuid.UUID, *ServErr)
	FindByUsername(username string) (dto.UserSingleResponse, *ServErr)
	Update(id uuid.UUID, request dto.UserUpdateRequest) *ServErr
	GetByUsername(username string) (*Users, *ServErr)
}

type MovieService interface {
	CreateMovie(request dto.MovieCreateRequest) (int64, *ServErr)
	FindById(id int64) (dto.MovieSingleResponse, *ServErr)
	FindAll() (dto.MovieListResponse, *ServErr)
	Update(id int64, request dto.MovieUpdateRequest) *ServErr
}

type LoginService interface {
	Login(w http.ResponseWriter, request dto.LoginRequest) (string, *ServErr)
	Logout(w http.ResponseWriter) *ServErr
}

type AuthService interface {
	Authenticate(usr *Users, pwd string) *ServErr
	IsPasswordsMatches(password string, hash []byte) bool
	GetJwtToken(usrId uuid.UUID, perms *dto.AllUserPermission) (string, *ServErr)
	VerifyJwt(req *http.Request) (*jwt.Token, *ServErr)
	HasPermission(tok *jwt.Token, permId int64, permLvl int16) *ServErr
	IsAuthenticated(tok *jwt.Token, usrId uuid.UUID) *ServErr
}

type PermissionService interface {
	GetAllUserPermissions(usr *Users) (*dto.AllUserPermission, *ServErr)
}

type RoleService interface {
	GetUserRolesIds(usrId uuid.UUID) (*[]int64, *ServErr)
}
