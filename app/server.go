package app

import (
	"fmt"
	"golang-echo-wasm-example/app/middlewares"
	"golang-echo-wasm-example/app/routers"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func RunServer(port int) {
	e := echo.New()

	middlewares.DefaultMiddlewares(e)

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("web/html/*.html")),
	}
	e.Renderer = renderer

	e.Static("/", "web")

	routers.MainRoute(e.Group(""))

	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
		e.Logger.Error(err)
		e.Logger.Panic("shuttig down the server")
	}
}
