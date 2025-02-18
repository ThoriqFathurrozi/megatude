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

func (e *EarthquakeHandler) GetLastEarthquake(ctx echo.Context) error {
	earthquake := earthquakeEntity.Earthquake{}

	if err := e.earthquakeRepo.FindLast(&earthquake); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, earthquake)
}

func (e *EarthquakeHandler) RefreshEarthquakeData(ctx echo.Context) error {
	cfg := configs.GetConfig()

	bmkg := bmkg.NewBMKG()

	autoBMKG, terkiniBMKG, dirasakanBMKG := bmkg.GetSourceData()

	autoEarthquake := earthquakeEntity.Earthquake{
		Datetime:    helpers.ParsingTime(autoBMKG.InfoGempa.Gampa.DateTime),
		Depth:       helpers.ParsingInt64(strings.Split(autoBMKG.InfoGempa.Gampa.Kedalaman, " ")[0]),
		Magnitude:   helpers.ParsingFloat64(autoBMKG.InfoGempa.Gampa.Magnitude),
		Location:    autoBMKG.InfoGempa.Gampa.Wilayah,
		Longitude:   autoBMKG.InfoGempa.Gampa.Bujur,
		Latitude:    autoBMKG.InfoGempa.Gampa.Lintang,
		Coordinates: autoBMKG.InfoGempa.Gampa.Coordinates,
	}

	earthquakes := []earthquakeEntity.Earthquake{}

	for _, v := range terkiniBMKG.InfoGempa.GempaTerkiniRes {
		earthquakes = append(earthquakes, earthquakeEntity.Earthquake{
			Datetime:    helpers.ParsingTime(v.DateTime),
			Depth:       helpers.ParsingInt64(strings.Split(v.Kedalaman, " ")[0]),
			Magnitude:   helpers.ParsingFloat64(v.Magnitude),
			Location:    v.Wilayah,
			Longitude:   v.Bujur,
			Latitude:    v.Lintang,
			Coordinates: v.Coordinates,
		})
	}

	for _, v := range dirasakanBMKG.InfoGempa.GempaDirasakanRes {
		earthquakes = append(earthquakes, earthquakeEntity.Earthquake{
			Datetime:    helpers.ParsingTime(v.DateTime),
			Depth:       helpers.ParsingInt64(strings.Split(v.Kedalaman, " ")[0]),
			Magnitude:   helpers.ParsingFloat64(v.Magnitude),
			Location:    v.Wilayah,
			Longitude:   v.Bujur,
			Latitude:    v.Lintang,
			Coordinates: v.Coordinates,
		})

	}

	earthquakes = append(earthquakes, autoEarthquake)

	rowAffected, err := e.earthquakeRepo.CreateAll(&earthquakes)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to refresh data", "error": err.Error()})
	}

	if rowAffected == 0 {
		return ctx.JSON(http.StatusOK, map[string]string{"message": "Data is up to date", "source": cfg.Resource.Url})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "Data refreshed", "source": cfg.Resource.Url})

}
