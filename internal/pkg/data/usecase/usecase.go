package usecase

import "github.com/keithzetterstrom/faf-user-service/internal/pkg/data"

type dataUsecase struct {
	dataRepo data.Repository
}

func New(dataRepo data.Repository) data.Usecase {
	return &dataUsecase{dataRepo: dataRepo}
}
