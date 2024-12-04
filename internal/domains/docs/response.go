package docs

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDocs(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
