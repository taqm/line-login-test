package handlers

import (
	"context"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func withOAuthHttpClientContext(src context.Context) context.Context {
	httpClient := &http.Client{Timeout: 2 * time.Second}
	return context.WithValue(src, oauth2.HTTPClient, httpClient)
}

func SetupHandler(e *echo.Echo) {
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.GET("/", indexHandler)
	e.GET("/login", loginHandler)
	e.GET("/auth/cb", authCallbackHandler)
}
