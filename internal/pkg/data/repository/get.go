package repository

import (
	"context"
	"github.com/keithzetterstrom/faf-user-service/internal/pkg/models"
	"github.com/keithzetterstrom/faf-user-service/internal/pkg/repository/queries"
	"github.com/pkg/errors"
)

func (d dataRepository) GetData(ctx context.Context, id int64) (*models.Data, error) {
	data, err := queries.GetDataByID(ctx, d.dbHandler, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get data")
	}

	return &models.Data{A: data}, nil
}
