package apis

import (
	"encoding/json"
	"io"
	"lfetchogger/apis/models"
	"net/http"
)

func (r Route) getLogs(c echo.Context) error {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Request().Body)

	var body models.LogRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := body.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	data, err := r.Logger.GetLogs(body.Search, body.From, body.To)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}
