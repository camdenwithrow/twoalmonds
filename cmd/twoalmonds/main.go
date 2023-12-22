package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/camdenwithrow/twoalmonds/template"
	"github.com/camdenwithrow/twoalmonds/views"
)

func main() {
	e := echo.New()

	e.Static("/dist", "dist")
	template.NewTemplateRenderer(e)

	e.GET("/", func(c echo.Context) error {
		component := views.Index()
		return template.AssertRender(c, http.StatusOK, component)
	})
	e.Logger.Fatal(e.Start(":8000"))
}
