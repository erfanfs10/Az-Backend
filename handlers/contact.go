package handlers

import (
	"net/http"

	"github.com/erfanfs10/Az-Backend/db"
	"github.com/erfanfs10/Az-Backend/models"
	"github.com/erfanfs10/Az-Backend/queries"
	"github.com/erfanfs10/Az-Backend/utils"
	"github.com/labstack/echo/v4"
)

func ContactCreate(c echo.Context) error {
	// bind data
	contactCreate := new(models.ContactCreate)
	if err := c.Bind(contactCreate); err != nil {
		return utils.HandleError(c, http.StatusBadRequest, err, "bad request")
	}
	// validate user input
	err := c.Validate(contactCreate)
	if err != nil {
		return utils.HandleError(c, http.StatusBadRequest,
			err, "invalid input")
	}
	// create the record
	_, err = db.DB.NamedExec(queries.ContactCreate, contactCreate)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "internal server error")
	}
	return c.JSON(http.StatusCreated,
		echo.Map{"message": "created"})
}
