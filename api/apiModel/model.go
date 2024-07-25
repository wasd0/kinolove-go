package apiModel

import (
	"fmt"
	"github.com/go-chi/render"
	"kinolove/internal/service"
	"net/http"
	"time"
)

type Response[T any] interface {
	Render(w http.ResponseWriter, r *http.Request) error
}

type RestResponse[T any] struct {
	Data *T `json:"data"`
}

func (rest *RestResponse[T]) Render(_ http.ResponseWriter, _ *http.Request) error {
	if rest.Data == nil {
		return fmt.Errorf("response data is nil")
	}

	return nil
}

type ErrResponse struct {
	Err error `json:"-"`

	Message string    `json:"message"`
	ErrCode int       `json:"code"`
	Time    time.Time `json:"time"`
	ErrDesc string    `json:"description,omitempty"`
}

func (e *ErrResponse) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.ErrCode)
	return nil
}
func NewErrRenderer(servErr *service.ServErr) render.Renderer {
	desc := ""
	if servErr.Err != nil {
		desc = servErr.Err.Error()
	}
	return &ErrResponse{
		Err:     servErr.Err,
		ErrCode: servErr.Code,
		Message: servErr.Msg,
		Time:    servErr.Time,
		ErrDesc: desc,
	}
}
