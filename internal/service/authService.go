package service

import (
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	. "kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/service/dto"
	"kinolove/internal/utils/mapper"
	"kinolove/pkg/constants"
	"kinolove/pkg/utils/crypt"
	"kinolove/pkg/utils/jwt"
	"net/http"
	"os"
	"time"
)

type AuthServiceImpl struct {
	jwtTok *jwt.Auth
}

func NewAuthService(jwtTok *jwt.Auth) *AuthServiceImpl {
	return &AuthServiceImpl{jwtTok: jwtTok}
}

func (a *AuthServiceImpl) Authenticate(usr *Users, pwd string) *ServErr {
	hash := usr.Password

	if !a.IsPasswordsMatches(pwd, hash) {
		return BadRequest(errors.New("Authentication error"), "Invalid username or password")
	}

	return nil
}

func (a *AuthServiceImpl) IsPasswordsMatches(password string, hash []byte) bool {
	return crypt.Matches([]byte(password), hash)
}

func (a *AuthServiceImpl) GetJwtToken(usrId uuid.UUID, perms *dto.AllUserPermission) (string, *ServErr) {
	tok := &jwt.Token{}

	exp := os.Getenv(constants.EnvExpIn)
	expIn, err := time.ParseDuration(exp)

	if err != nil {
		return "", InternalError(err)
	}

	tok.Sub = usrId

	userPerms, rolePerms := mapper.PermissionToJwt(perms)

	tok.UserPerms = *userPerms
	tok.RolePerms = *rolePerms
	tok.ExpIn = expIn

	tokenStr, jwtErr := a.jwtTok.Encode(tok)

	if jwtErr != nil {
		return "", InternalError(jwtErr)
	}

	return tokenStr, nil
}

func (a *AuthServiceImpl) VerifyJwt(req *http.Request) *ServErr {
	token := jwtauth.TokenFromCookie(req)

	err := a.jwtTok.Verify(token)

	if err != nil {
		return Unauthorized(err)
	}

	return nil
}
