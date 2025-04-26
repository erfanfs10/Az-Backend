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

func CommentList(c echo.Context) error {
	// create model and query
	comments := []models.Comment{}
	// get the data from db
	err := db.DB.Select(&comments, queries.CommentList)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound,
				err, "comment not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "internal server error")
	}
	// create response
	var response struct {
		Comments []models.Comment `json:"comments"`
	}
	response.Comments = comments
	// return the response
	return c.JSON(http.StatusOK, response)
}
