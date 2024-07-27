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

func RenderError(w http.ResponseWriter, r *http.Request, servErr *service.ServErr, log logger.Common) {
	renderErr := render.Render(w, r, apiModel.NewErrRenderer(servErr))
	if renderErr != nil {
		log.Fatal(renderErr, "rendering error")
	}
}
