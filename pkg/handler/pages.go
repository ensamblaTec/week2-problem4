package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RunLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", GetProducts())
}
