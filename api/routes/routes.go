package router

import (
	"github.com/rsanantmishra/metubeplus/api/handler"

	"github.com/gofiber/fiber/v2" // swagger handler
	"github.com/gofiber/swagger"
	_ "github.com/rsanantmishra/metubeplus/cmd/docs"
)

func SetupRoutes(app *fiber.App) {

	//Middlewares
	//api := app.Group("/api", logger.New())
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/", handler.Hello)
	app.Get("/Mello", handler.Mello)

	//routes Videos, Files, Tags
}
