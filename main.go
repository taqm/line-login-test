package main

import (
	"encoding/gob"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/taqm/line-login-test/domains"
	"github.com/taqm/line-login-test/handlers"
)

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	gob.Register(domains.User{})

	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	handlers.SetupHandler(e)

	e.Logger.Fatal(e.Start(":" + port))
}
