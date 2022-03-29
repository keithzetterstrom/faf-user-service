package usecase

import (
	"context"

	"github.com/keithzetterstrom/faf-user-service/internal/pkg/models"
)

func (d dataUsecase) GetData(ctx context.Context, id int64) (*models.Data, error) {
	data, err := d.dataRepo.GetData(ctx, id)
	if err != nil {
		return nil, err
	}

	data.ID = id

	return data, nil
}
