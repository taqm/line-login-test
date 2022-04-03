package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taqm/line-login-test/domains"
	"github.com/taqm/line-login-test/line"
	"github.com/taqm/line-login-test/mysession"
)

func authCallbackHandler(c echo.Context) error {
	ctx := withOAuthHttpClientContext(c.Request().Context())

	code := c.QueryParam("code")
	client, err := line.NewLineApiClientByCode(ctx, code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	profile, err := client.GetProfile()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	user := domains.User{
		ID:          profile.UserID,
		DisplayName: profile.DisplayName,
		PictureURL:  profile.PictureURL,
	}
	sess := mysession.GetFromContext(c)
	err = sess.SaveUser(user, c)
	if err != nil {
		c.Error(err)
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/")
}
