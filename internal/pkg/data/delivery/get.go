package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	outmodel "github.com/keithzetterstrom/faf-user-service/cmd/handlers/models"
)

func (d dataDelivery) GetData(c echo.Context) error {
	id := c.Param("id")

	parseID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	data, err := d.dataUsecase.GetData(c.Request().Context(), parseID)
	if err != nil {
		return err
	}

	outData := outmodel.ModelToData(data)

	return c.JSON(http.StatusOK, outData)
}
