package logger

import (
	"context"
	"fmt"
	contextlib "github.com/keithzetterstrom/faf-user-service/utils/context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type logger struct {
	logger *zap.Logger
}

type Logger interface {
	Error(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	With(fields ...zap.Field) Logger
	WithContext(ctx context.Context) Logger
}

func NewLogger(serviceName string) (Logger, error) {
	zLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	zLogger = zLogger.With(zap.String("service", serviceName))

	return &logger{logger: zLogger}, nil
}

func (l logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l logger) With(fields ...zap.Field) Logger {
	l.logger = l.logger.With(fields...)
	return l
}

func (l logger) WithContext(ctx context.Context) Logger {
	reqID := contextlib.GetRequestID(ctx)
	if reqID == nil {
		return l
	}
	l.logger = l.logger.With(zap.String("request_id", *reqID))
	return l
}

func EchoRequestLogger(log Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			fields := []zapcore.Field{
				zap.String("remote_ip", c.RealIP()),
				zap.String("latency", time.Since(start).String()),
				zap.String("host", req.Host),
				zap.String("request", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Int("status", res.Status),
				zap.Int64("size", res.Size),
				zap.String("user_agent", req.UserAgent()),
			}

			id := req.Header.Get(contextlib.RequestIDHeader)
			if id == "" {
				id = res.Header().Get(contextlib.RequestIDHeader)
			}

			if id != "" {
				fields = append(fields, zap.String("request_id", id))
			}

			n := res.Status
			switch {
			case n >= 500:
				log.With(zap.Error(err)).Error("Server error", fields...)
			case n >= 400:
				log.With(zap.Error(err)).Warn("Client error", fields...)
			case n >= 300:
				log.Info("Redirection", fields...)
			default:
				log.Info("Success", fields...)
			}

			return nil
		}
	}
}
