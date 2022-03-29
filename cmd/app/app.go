package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/keithzetterstrom/faf-user-service/cmd/handlers"
	"github.com/keithzetterstrom/faf-user-service/internal/pkg/data/delivery"
	"github.com/keithzetterstrom/faf-user-service/internal/pkg/data/repository"
	"github.com/keithzetterstrom/faf-user-service/internal/pkg/data/usecase"
	db "github.com/keithzetterstrom/faf-user-service/internal/pkg/repository/postgres"
	faflogger "github.com/keithzetterstrom/faf-user-service/utils/logger"
	"github.com/keithzetterstrom/faf-user-service/utils/must"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type baseAPI struct {
	echoService *echo.Echo
	logger      faflogger.Logger
	dbHandler   db.Handler
	cfg         Config
}

type Base interface {
	ServerStart()
	Close()
}

func Start() Base {
	var cfg Config

	err := NewConfig(&cfg)
	must.Must(err)

	logger, err := faflogger.NewLogger("faf-user-service")
	must.Must(err)

	dbContext := context.Background()
	dbHandler, err := db.New(dbContext, cfg.PostgresConfig.Config, logger)
	must.Must(err)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	base := baseAPI{
		echoService: e,
		dbHandler:   dbHandler,
		logger:      logger,
		cfg:         cfg,
	}

	base.ServiceAPI()

	return &base
}

func (b *baseAPI) ServiceAPI() {
	dataRepo := repository.New(b.dbHandler)

	dataUsecase := usecase.New(dataRepo)

	dataDelivery := delivery.New(dataUsecase)

	handlers.Router(b.echoService, dataDelivery, b.logger)
}

func (b *baseAPI) Close() {
	b.echoService.Close()
	b.dbHandler.Close()
	b.logger.Info("stop server")
}

func (b *baseAPI) ServerStart() {
	b.logger.Info(
		"start server",
		zap.String("address", fmt.Sprintf("%s:%s", b.cfg.ServiceConfig.Host, b.cfg.ServiceConfig.Port)),
	)

	if err := b.echoService.Start(fmt.Sprintf(
		"%s:%s",
		b.cfg.ServiceConfig.Host,
		b.cfg.ServiceConfig.Port,
	)); err != nil && err != http.ErrServerClosed {
		b.logger.Error("shutting down the server")
	}
}
