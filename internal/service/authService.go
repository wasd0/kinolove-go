package service

import (
	"github.com/go-chi/jwtauth"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/pkg/errors"
	"kinolove/internal/consts/perms"
	. "kinolove/internal/entity/.gen/kinolove/public/model"
	"kinolove/internal/service/dto"
	"kinolove/internal/utils/mapper"
	"kinolove/pkg/constants"
	"kinolove/pkg/utils/crypt"
	"kinolove/pkg/utils/jwtUtils"
	"net/http"
	"os"
	"strconv"
	"time"
)

type AuthServiceImpl struct {
	jwtTok *jwtUtils.Auth
}

func NewAuthService(jwtTok *jwtUtils.Auth) *AuthServiceImpl {
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
	tok := &jwtUtils.Token{}

	exp := os.Getenv(constants.EnvExpIn)
	expIn, err := time.ParseDuration(exp)

	if err != nil {
		return "", InternalError(err)
	}

	userPerms, rolePerms := mapper.PermissionToMaps(perms)

	tok.Sub = usrId.String()
	tok.UserPerms = *userPerms
	tok.RolePerms = *rolePerms
	tok.ExpIn = expIn

	tokenStr, jwtErr := a.jwtTok.Encode(tok)

	if jwtErr != nil {
		return "", InternalError(jwtErr)
	}

	return tokenStr, nil
}

func (a *AuthServiceImpl) VerifyJwt(req *http.Request) (*jwt.Token, *ServErr) {
	token, _, err := jwtauth.FromContext(req.Context())

	if err != nil {
		return nil, Unauthorized(err)
	}

	if token == nil || jwt.Validate(token) != nil {
		return nil, Unauthorized(err)
	}

	return &token, nil
}

func (a *AuthServiceImpl) HasPermission(tok *jwt.Token, permId int64, permLevel int16) *ServErr {
	if tok == nil {
		return Forbidden(errors.New("Token is nil"))
	}

	var usrPermLvl interface{} = nil

	permStrId := strconv.Itoa(int(permId))

	if usrPerms, isOK := (*tok).Get(perms.UserPerms); isOK {
		permsMap := usrPerms.(map[string]interface{})
		if permsMap != nil {
			if perm, hasPerm := permsMap[permStrId]; hasPerm {
				usrPermLvl = perm
			}
		}
	}

	if rolePerms, isOK := (*tok).Get(perms.RolePerms); isOK {
		permsMap := rolePerms.(map[string]interface{})
		if permsMap != nil {
			if perm, hasPerm := permsMap[permStrId]; hasPerm {
				usrPermLvl = perm
			}
		}
	}

	if val, isOk := usrPermLvl.(float64); !isOk || int16(val) < permLevel {
		return Forbidden(errors.New("Forbidden"))
	}

	return nil
}

func (a *AuthServiceImpl) IsAuthenticated(tok *jwt.Token, usrId uuid.UUID) *ServErr {
	if usrId == uuid.Nil || len((*tok).Subject()) == 0 {
		return Unauthorized(errors.New("Invalid user"))
	}

	if err := uuid.Validate((*tok).Subject()); err != nil {
		return Unauthorized(err)
	}

	if usrId.String() != (*tok).Subject() {
		return Unauthorized(errors.New("Invalid user"))
	}

	return nil
}
