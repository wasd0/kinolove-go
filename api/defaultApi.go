package api

import (
	"kinolove/internal/service"
	"kinolove/pkg/logger"
	"net/http"
)

type DefaultApi struct {
	log logger.Common
}

func NewDefaultApi(log logger.Common) *DefaultApi {
	return &DefaultApi{log: log}
}

func (api *DefaultApi) NotFound(w http.ResponseWriter, r *http.Request) {
	renderError(w, r, service.NotFound("Page not found"), api.log)
	return
}

func (api *DefaultApi) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	renderError(w, r, service.MethodNotAllowed("Method is not allowed"), api.log)
	return
}
