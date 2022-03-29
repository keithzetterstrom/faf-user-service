package middleware

import (
	"fmt"
	"math/rand"

	"github.com/labstack/echo/v4"
	contextlib "github.com/keithzetterstrom/faf-user-service/utils/context"
)

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()

		requestID := req.Header.Get(contextlib.RequestIDHeader)
		if requestID == "" {
			requestID = fmt.Sprintf("%016x", rand.Int()) // nolint:gosec
		}

		c.Response().Header().Set(contextlib.RequestIDHeader, requestID)

		ctx := contextlib.SetRequestID(c.Request().Context(), requestID)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
