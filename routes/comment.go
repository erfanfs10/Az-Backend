package routes

import (
	"github.com/erfanfs10/Az-Backend/handlers"
	"github.com/labstack/echo/v4"
)

func CommentRoutes(g *echo.Group) {
	g.GET("", handlers.CommentList)
}
