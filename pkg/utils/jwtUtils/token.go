package jwtUtils

import (
	"time"
)

type Token struct {
	Sub       string          `json:"sub"`
	UserPerms map[int64]int16 `json:"user_permissions"`
	RolePerms map[int64]int16 `json:"role_permissions"`
	ExpIn     time.Duration   `json:"expIn"`
}
