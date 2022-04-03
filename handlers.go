package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func setupHandler(e *echo.Echo) {
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"message": "Hello World",
		})
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "https://google.com")
	})
}
