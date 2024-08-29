package api

import (
	"kinolove/internal/service"
	"net/http"
)

type DefaultApi struct {
}

func NewDefaultApi() *DefaultApi {
	return &DefaultApi{}
}

func (api *DefaultApi) NotFound(w http.ResponseWriter, r *http.Request) {
	RenderError(w, r, service.NotFound("Page not found"))
	return
}

func (api *DefaultApi) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	RenderError(w, r, service.MethodNotAllowed("Method is not allowed"))
	return
}
