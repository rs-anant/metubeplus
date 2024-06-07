package main

import (
	"log"

	"github.com/rsanantmishra/metubeplus/api/middleware"
	r "github.com/rsanantmishra/metubeplus/api/routes"
	c "github.com/rsanantmishra/metubeplus/config"

	_ "github.com/rsanantmishra/metubeplus/cmd/docs"

	"github.com/gofiber/fiber/v2"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "MeTube+",
	})

	middleware.SetupMiddleware(app)
	r.SetupRoutes(app)

	log.Fatal(app.Listen(":" + c.Config("APP_PORT")))
}
