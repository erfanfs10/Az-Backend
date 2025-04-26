package routes

import (
	"github.com/erfanfs10/Az-Backend/handlers"
	"github.com/labstack/echo/v4"
)

func ServiceRoutes(g *echo.Group) {
	g.GET("", handlers.ServiceList)
}
