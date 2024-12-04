package docs

import (
	"net/http"

	"github.com/ThoriqFathurrozi/megatude/configs"
	"github.com/labstack/echo/v4"
)

// Handler struct
type DocsHandler struct {
	version string
}

func NewDocsHandler() *DocsHandler {
	cfg := configs.GetConfig()

	return &DocsHandler{
		version: cfg.App.Version,
	}
}

// GetDocs function
func (h *DocsHandler) GetDocs(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"version": h.version})
}
