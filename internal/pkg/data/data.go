package data

import (
	"context"

	"github.com/keithzetterstrom/faf-user-service/internal/pkg/models"
	"github.com/labstack/echo/v4"
)

type Delivery interface {
	GetData(c echo.Context) error
}

type Repository interface {
	GetData(ctx context.Context, id int64) (*models.Data, error)
}

type Usecase interface {
	GetData(ctx context.Context, id int64) (*models.Data, error)
}
