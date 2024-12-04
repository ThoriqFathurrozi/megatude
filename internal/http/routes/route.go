package routes

import (
	"github.com/labstack/echo/v4"

	docsDomain "github.com/ThoriqFathurrozi/megatude/internal/domains/docs"
)

type Route struct {
	App         *echo.Echo
	DocsHandler *docsDomain.DocsHandler
}

func (r *Route) InitializeV1() {
	api := r.App.Group("/api")
	v1 := api.Group("/v1")

	r.initializeRoutes(v1)
}

func (r *Route) initializeRoutes(router *echo.Group) {
	docs := router.Group("/docs")

	docs.GET("", r.DocsHandler.GetDocs)

}
