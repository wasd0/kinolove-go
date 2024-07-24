package service

import (
	"time"
)

type ServErr struct {
	Code int
	Msg  string
	Time time.Time
	Err  error
}

func BadRequest(err error, msg string) *ServErr {

	return &ServErr{
		Code: 400,
		Msg:  msg,
		Time: time.Now().UTC(),
		Err:  err,
	}
}

func InternalError(err error) *ServErr {
	return &ServErr{
		Code: 500,
		Msg:  "Internal Server Error",
		Time: time.Now().UTC(),
		Err:  err,
	}
}

func NotFound(msg string) *ServErr {
	return &ServErr{
		Code: 404,
		Msg:  msg,
		Time: time.Now().UTC(),
	}
}
