package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func Router() *echo.Echo {
	r := echo.New()
	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	return r
}
