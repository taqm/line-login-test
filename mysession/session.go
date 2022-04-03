package mysession

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/taqm/line-login-test/domains"
)

type MySession struct {
	sess *sessions.Session
}

func (s MySession) GetUser() domains.User {
	var user domains.User
	if s.sess.Values["User"] != nil {
		user = s.sess.Values["User"].(domains.User)
	}
	return user
}

func (s MySession) SaveUser(user domains.User, c echo.Context) error {
	s.sess.Values["User"] = user
	return s.sess.Save(c.Request(), c.Response())
}

type SessionData struct {
	User domains.User
}

func GetFromContext(c echo.Context) MySession {
	sess, _ := session.Get("session", c)
	return MySession{sess}
}
