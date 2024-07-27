package jwt

import (
	"github.com/google/uuid"
	"time"
)

type Token struct {
	Sub       uuid.UUID            `json:"sub"`
	UserPerms map[int64]Permission `json:"user_permissions"`
	RolePerms map[int64]Permission `json:"role_permissions"`
	ExpIn     time.Duration        `json:"expIn"`
}

type Permission struct {
	TargetLvl int16 `json:"target_lvl"`
	GlobalLvl int16 `json:"global_lvl"`
}
