package core

import (
	"fmt"

	"github.com/ThoriqFathurrozi/megatude/configs"
	"github.com/ThoriqFathurrozi/megatude/internal/domains/docs"
	"github.com/ThoriqFathurrozi/megatude/internal/http/routes"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Megatude struct {
	Config *configs.Config
	App    *echo.Echo
	DB     *gorm.DB
}

func Init(megatude *Megatude) {
	docsHandler := docs.NewDocsHandler()

	route := routes.Route{
		App:         megatude.App,
		DocsHandler: docsHandler,
	}

	route.InitializeV1()

}

func (a *Megatude) Start() {
	addr := fmt.Sprintf(":%d", a.Config.App.Port)

	a.App.Logger.Fatal(a.App.Start(addr))
}
