package core

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ThoriqFathurrozi/megatude/configs"
	"github.com/ThoriqFathurrozi/megatude/internal/domains/docs"
	"github.com/ThoriqFathurrozi/megatude/internal/domains/earthquake"
	earthquakeRepository "github.com/ThoriqFathurrozi/megatude/internal/domains/earthquake/repository"
	"github.com/ThoriqFathurrozi/megatude/internal/http/routes"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Megatude struct {
	Config *configs.Config
	App    *echo.Echo
	DB     *gorm.DB
	Corn   *cron.Cron
}

func Init(megatude *Megatude) {
	docsHandler := docs.NewDocsHandler()

	earthquakeRepo := earthquakeRepository.NewEarthquakeRepository(megatude.DB)
	earthquakeHandler := earthquake.NewEarthquakeHandler(earthquakeRepo)

	route := routes.Route{
		App:         megatude.App,
		DocsHandler: docsHandler,
		Earthquake:  earthquakeHandler,
	}

	route.InitializeV1()

}

func (a *Megatude) Start() {
	addr := fmt.Sprintf(":%d", a.Config.App.Port)
	schedule := fmt.Sprintf("*/%v * * * *", a.Config.Cron.Interval)

	a.Corn.AddJob(schedule, cron.FuncJob(func() {
		fmt.Println("Running cron job")
		go func() {
			res, err := http.Get("http://localhost:5555/api/v1/earthquake/refresh")
			if err != nil {
				fmt.Println(err)
			}
			defer res.Body.Close()

			resData, err := io.ReadAll(res.Body)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(resData))
			fmt.Println("Cron job done")
		}()
	}))

	a.Corn.Start()

	a.App.Logger.Fatal(a.App.Start(addr))

}
