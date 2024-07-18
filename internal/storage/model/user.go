package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id          uuid.UUID  `ksql:"id"`
	Username    string     `ksql:"username"`
	Password    string     `ksql:"password"`
	IsActive    bool       `ksql:"is_active"`
	DateReg     time.Time  `ksql:"date_reg,timeNowUTC/skipUpdates"`
	DatePassUpd *time.Time `ksql:"date_pass_upd"`
}
