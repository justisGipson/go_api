package main

import (
	"github.com/CodeliciousProduct/bluebird/internal/configs"
	"github.com/CodeliciousProduct/bluebird/internal/middleware"
	"github.com/CodeliciousProduct/bluebird/internal/routes"
	"github.com/CodeliciousProduct/bluebird/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"

	_ "github.com/CodeliciousProduct/bluebird/third_party" // swagger docs
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// define fiber config
	config := configs.FiberConfig()
	// new fiber app
	app := fiber.New(config)
	godotenv.Load()
	// https headers with fiber/helmet
	app.Use(helmet.New())
	// register fiber's middleware
	middleware.FiberMiddleware(app)
	// routes
	routes.SwaggerRoute(app)  // API docs
	routes.RouteNotFound(app) // 404 page
	routes.PrivateRoutes(app) // all internal routes
	routes.PublicRoutes(app)  // public routes - just sign in
	// start up, now comes with graceful shutdown
	utils.StartServer(app)
}
