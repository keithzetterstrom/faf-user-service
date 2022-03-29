package queries

import (
	"context"

	db "github.com/keithzetterstrom/faf-user-service/internal/pkg/repository/postgres"
)

func GetDataByID(ctx context.Context, dbHandler db.Handler, id int64) (string, error) {
	return "some data", nil
}
