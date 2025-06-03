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

func PortfolioList(c echo.Context) error {
	// create model and query
	portfolios := []models.Portfolio{}
	// get the data from db
	err := db.DB.Select(&portfolios, queries.PortfolioList)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound,
				err, "portfolio not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "internal server error")
	}
	// create response
	var response struct {
		Portfolios []models.Portfolio `json:"portfolios"`
	}
	response.Portfolios = portfolios
	// return the response
	return c.JSON(http.StatusOK, response)
}
