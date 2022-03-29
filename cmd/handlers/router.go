package handlers

import (
	"github.com/keithzetterstrom/faf-user-service/internal/pkg/data"
	"github.com/keithzetterstrom/faf-user-service/internal/pkg/middleware"
	faflogger "github.com/keithzetterstrom/faf-user-service/utils/logger"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Router(
	e *echo.Echo,
	dataDelivery data.Delivery,
	logger faflogger.Logger,
) {
	api := e.Group("")

	h := promhttp.Handler()
	e.Any("/metrics", echo.WrapHandler(h))

	api.Use(middleware.RequestIDMiddleware, faflogger.EchoRequestLogger(logger))

	api.GET("/example/:id", dataDelivery.GetData)
}
