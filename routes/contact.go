package routes

import (
	"github.com/erfanfs10/Az-Backend/handlers"
	"github.com/labstack/echo/v4"
)

func ContactRoutes(g *echo.Group) {
	g.POST("", handlers.ContactCreate)
}
