package dto

import "time"

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSingleResponse struct {
	Username    string     `json:"username"`
	IsActive    bool       `json:"isActive"`
	DateReg     *time.Time `json:"dateReg"`
	DatePassUpd *time.Time `json:"datePassUpd"`
}

type UserUpdateRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
