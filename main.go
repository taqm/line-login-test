package main

import (
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	e := echo.New()
	setupHandler(e)

	e.Logger.Fatal(e.Start(":" + port))
}
