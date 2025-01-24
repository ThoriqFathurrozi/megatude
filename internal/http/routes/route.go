package routes

import (
	"github.com/labstack/echo/v4"

	docsDomain "github.com/ThoriqFathurrozi/megatude/internal/domains/docs"
	quakeDomain "github.com/ThoriqFathurrozi/megatude/internal/domains/earthquake"
)

type Route struct {
	App         *echo.Echo
	DocsHandler *docsDomain.DocsHandler
	Earthquake  *quakeDomain.EarthquakeHandler
}

func (r *Route) InitializeV1() {
	api := r.App.Group("/api")
	v1 := api.Group("/v1")

	r.initializeRoutes(v1)
}

func (r *Route) initializeRoutes(router *echo.Group) {
	docs := router.Group("/docs")

	docs.GET("", r.DocsHandler.GetDocs)

	earthquake := router.Group("/earthquake")

	earthquake.GET("", r.Earthquake.GetEarhquake)
	earthquake.GET("/source", r.Earthquake.GetSourceData)
	earthquake.GET("/refresh", r.Earthquake.RefreshEarthquakeData)
	earthquake.GET("/last", r.Earthquake.GetLastEarthquake)
}
