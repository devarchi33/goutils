package echomiddleware

import (
	"context"

	"github.com/devarchi33/goutils/ctxbase"
	"github.com/labstack/echo/v4"
)

func ContextBase() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			rid := req.Header.Get(echo.HeaderXRequestID)
			if rid == "" {
				rid = ctxbase.NewID()
			}
			c.Response().Header().Set(echo.HeaderXRequestID, rid)

			cb := ctxbase.ContextBase{
				RequestID: rid,
				ActionID:  ctxbase.NewID(),
			}
			c.SetRequest(req.WithContext(context.WithValue(req.Context(), ctxbase.ContextBaseName, &cb)))
			return next(c)
		}
	}
}
