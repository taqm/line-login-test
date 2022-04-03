package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taqm/line-login-test/line"
)

func loginHandler(c echo.Context) error {
	url := line.AuthCodeURL()
	return c.Redirect(http.StatusSeeOther, url)
}
