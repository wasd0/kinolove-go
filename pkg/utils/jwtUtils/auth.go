package jwtUtils

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/jwtauth"
	"kinolove/internal/consts/perms"
	"kinolove/pkg/constants"
	"os"
	"time"
)

type Auth struct {
	jwt *jwtauth.JWTAuth
}

func NewJwtAuth() *Auth {
	secret := os.Getenv(constants.EnvJwtSecret)
	alg := os.Getenv(constants.EnvJwtAlg)
	token := jwtauth.New(alg, []byte(secret), nil)
	return &Auth{jwt: token}
}

func (a *Auth) GetJwt() *jwtauth.JWTAuth {
	return a.jwt
}

func (a *Auth) Encode(token *Token) (string, error) {
	claims := make(map[string]interface{}, 3)
	claims[perms.UserPerms] = token.UserPerms
	claims[perms.RolePerms] = token.RolePerms
	claims[perms.Sub] = token.Sub
	jwtauth.SetExpiryIn(claims, token.ExpIn)

	_, tokString, err := a.jwt.Encode(claims)
	if err != nil {
		return "", fmt.Errorf("encode jwtUtils error: %v", err)
	}

	return tokString, nil
}

func (a *Auth) Decode(token string) (*Token, error) {
	jwtTok, err := a.jwt.Decode(token)
	res := &Token{}

	if err != nil {
		return res, fmt.Errorf("JWT docode fail: %v", err)
	}

	if userPerms, isOk := jwtTok.Get("user_permissions"); isOk {
		if mErr := json.Unmarshal([]byte(userPerms.(string)), &res.UserPerms); mErr != nil {
			return res, fmt.Errorf("user permissions docode fail: %v", mErr)
		}
	}

	if rolePerms, isOk := jwtTok.Get("role_permissions"); isOk {
		if mErr := json.Unmarshal([]byte(rolePerms.(string)), &res.RolePerms); mErr != nil {
			return res, fmt.Errorf("permissions docode fail: %v", mErr)
		}
	}

	if expIn, isOk := jwtTok.Get("expIn"); isOk {
		res.ExpIn = expIn.(time.Duration)
	}

	sub := jwtTok.Subject()
	res.Sub = sub

	return res, nil
}
