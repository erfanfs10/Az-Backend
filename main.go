package main

import (
	"github.com/erfanfs10/Az-Backend/db"
	"github.com/erfanfs10/Az-Backend/middlewares"
	"github.com/erfanfs10/Az-Backend/routes"
	"github.com/erfanfs10/Az-Backend/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"strings"
)

/*
init func runs before main
load env , connect to database and create tables
before program starts
*/
func init() {
	utils.LoadEnv()
	db.ConnectDb()
	db.CreateTables()
}

func main() {
	// create echo instance
	e := echo.New()
	// asign custom validator
	e.Validator = utils.CreateCustomValidator()
	// close connection when stop the program
	defer db.DB.Close()
	// using CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(utils.GetEnv("ALLOWED_CORS"), ","),
		AllowMethods: []string{echo.GET, echo.POST,
			echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin,
			echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAuthorization},
	}))
	// using custom middlewares
	e.Use(middlewares.CustomLogger())
	e.Use(middlewares.SeparateLogs())
	e.Use(middleware.Recover())
	// serving static files in dev mode
	if os.Args[len(os.Args)-1] == "dev" {
		e.Static("media", "../")
	}
	// root route
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Az-Lab-Backend")
	})
	// define routes
	routes.ServiceRoutes(e.Group("api/services/"))
	routes.CommentRoutes(e.Group("api/comments/"))
	routes.TeamRoutes(e.Group("api/team/"))
	routes.ContactRoutes(e.Group("api/contact/"))
	// serve the program
	e.Logger.Fatal(e.Start(":8000"))
}
