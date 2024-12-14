package earthquake

import (
	"net/http"
	"strings"

	"github.com/ThoriqFathurrozi/megatude/configs"
	earthquakeEntity "github.com/ThoriqFathurrozi/megatude/internal/domains/earthquake/entity"
	earthquakeRepository "github.com/ThoriqFathurrozi/megatude/internal/domains/earthquake/repository"
	"github.com/ThoriqFathurrozi/megatude/internal/helpers"
	"github.com/ThoriqFathurrozi/megatude/third_party/bmkg"
	"github.com/labstack/echo/v4"
)

type EarthquakeHandler struct {
	earthquakeRepo *earthquakeRepository.EarthquakeRepository
	cfg            *configs.Config
}

func NewEarthquakeHandler(earthquakeRepo *earthquakeRepository.EarthquakeRepository) *EarthquakeHandler {
	cfg := configs.GetConfig()
	return &EarthquakeHandler{
		earthquakeRepo: earthquakeRepo,
		cfg:            cfg,
	}
}

func (e *EarthquakeHandler) GetEarhquake(ctx echo.Context) error {
	earthquakes := []earthquakeEntity.Earthquake{}

	if err := e.earthquakeRepo.FindAll(&earthquakes); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if len(earthquakes) == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{"message": "No data found"})
	}

	return ctx.JSON(http.StatusOK, earthquakes)
}

func (e *EarthquakeHandler) GetSourceData(ctx echo.Context) error {
	cfg := configs.GetConfig()
	url := cfg.Resource.Url

	return ctx.JSON(http.StatusOK, map[string]string{"source": url})
}

func (e *EarthquakeHandler) RefreshEarthquakeData(ctx echo.Context) error {

	cfg := configs.GetConfig()

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Data refreshed", "url": cfg.Resource.Url})

}

func (e *EarthquakeHandler) GetAutoEarthquake(ctx echo.Context) error {
	bmkg := bmkg.NewBMKG()

	autoBMKG := bmkg.GetSourceData()

	autoEarthquake := earthquakeEntity.Earthquake{
		Datetime:    autoBMKG.InfoGempa.Gampa.DateTime,
		Depth:       helpers.ParsingInt64(strings.Split(autoBMKG.InfoGempa.Gampa.Kedalaman, " ")[0]),
		Magnitude:   helpers.ParsingFloat64(autoBMKG.InfoGempa.Gampa.Magnitude),
		Location:    autoBMKG.InfoGempa.Gampa.Wilayah,
		Longitude:   autoBMKG.InfoGempa.Gampa.Bujur,
		Latitude:    autoBMKG.InfoGempa.Gampa.Lintang,
		Coordinates: autoBMKG.InfoGempa.Gampa.Coordinates,
	}

	if err := e.earthquakeRepo.Create(&autoEarthquake); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	

	return ctx.JSON(http.StatusOK, autoEarthquake)

}
