package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/erfanfs10/Az-Backend/db"
	"github.com/erfanfs10/Az-Backend/models"
	"github.com/erfanfs10/Az-Backend/queries"
	"github.com/erfanfs10/Az-Backend/utils"
	"github.com/labstack/echo/v4"
)

func ServiceList(c echo.Context) error {
	// create model and query
	servieces := []models.Service{}
	// get the data from db
	err := db.DB.Select(&servieces, queries.ServiceList)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound,
				err, "service not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "internal server error")
	}
	// create response
	var response struct {
		Servieces []models.Service `json:"services"`
	}
	response.Servieces = servieces
	// return the response
	return c.JSON(http.StatusOK, response)
}
