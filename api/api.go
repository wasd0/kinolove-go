package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"kinolove/api/apiModel"
	"kinolove/internal/service"
	"kinolove/pkg/logger"
	"net/http"
)

type ChiApi interface {
	Register() (string, func(router chi.Router))
	Handle(router chi.Router)
}

func renderError(w http.ResponseWriter, r *http.Request, servErr *service.ServErr, log logger.Common) {
	renderErr := render.Render(w, r, apiModel.NewErrRenderer(servErr))
	if renderErr != nil {
		log.Fatal(renderErr, "rendering error")
	}
}

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
