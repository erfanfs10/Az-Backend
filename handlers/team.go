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

func TeamList(c echo.Context) error {
	// create model and query
	team := []models.Team{}
	// get the data from db
	err := db.DB.Select(&team, queries.TeamList)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound,
				err, "team not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "internal server error")
	}
	// create response
	var response struct {
		Team []models.Team `json:"team"`
	}
	response.Team = team
	// return the response
	return c.JSON(http.StatusOK, response)
}
