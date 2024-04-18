package handlers

import (
	"github.com/labstack/echo/v4"

	"fiveloved/views"
)

func Index(c echo.Context) error {
	return renderer(c, views.Index())
}
