package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taqm/line-login-test/mysession"
)

func indexHandler(c echo.Context) error {
	sess := mysession.GetFromContext(c)
	user := sess.GetUser()
	if user.IsZeroValue() {
		return c.Render(http.StatusOK, "not_logged_in.html", nil)
	}
	return c.Render(http.StatusOK, "logged_in.html", user)
}
