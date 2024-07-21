package user

import "time"

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SingleResponse struct {
	Username    string     `json:"username"`
	IsActive    bool       `json:"isActive"`
	DateReg     *time.Time `json:"dateReg"`
	DatePassUpd *time.Time `json:"datePassUpd"`
}
